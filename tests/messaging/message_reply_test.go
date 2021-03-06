package messaging

import (
	"fmt"
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"

	"github.com/cortezaproject/corteza-server/tests/helpers"
)

func TestMessagesReply(t *testing.T) {
	h := newHelper(t)
	msg := h.makeMessage("old", h.repoMakePublicCh(), h.cUser)

	reply := func() uint64 {
		rval := struct {
			Response struct {
				ID uint64 `json:"messageID,string"`
			}
		}{}
		h.apiInit().
			Post(fmt.Sprintf("/channels/%d/messages/%d/replies", msg.ChannelID, msg.ID)).
			Header("Accept", "application/json").
			JSON(`{"message":"new reply"}`).
			Expect(t).
			Status(http.StatusOK).
			Assert(helpers.AssertNoErrors).
			Assert(jsonpath.Present(`$.response.messageID`)).
			Assert(jsonpath.Present(`$.response.replyTo`)).
			Assert(jsonpath.Equal(`$.response.message`, `new reply`)).
			End().
			JSON(&rval)

		r, _ := h.lookupMessageByID(rval.Response.ID)
		h.a.Equal(`new reply`, r.Message)
		h.a.Equal(msg.ID, r.ReplyTo)
		return rval.Response.ID
	}

	reply1ID := reply()
	reply2ID := reply()
	reply3ID := reply()

	_, _, _ = reply1ID, reply2ID, reply3ID

	msg, _ = h.lookupMessageByID(msg.ID)
	h.a.Equal(msg.Replies, uint(3))

	h.apiInit().
		Get("/search/threads").
		Header("Accept", "application/json").
		Query("channelID", fmt.Sprintf("%d", msg.ChannelID)).
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		Assert(jsonpath.Len(`$.response`, 4)). // 3 replies + original msg
		End()

	// Remove one of the replies
	h.apiInit().
		Delete(fmt.Sprintf("/channels/%d/messages/%d", msg.ChannelID, reply2ID)).
		Header("Accept", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		End()

	msg, _ = h.lookupMessageByID(msg.ID)
	h.a.Equal(msg.Replies, uint(2))

}
