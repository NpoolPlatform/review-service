package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/review-service"
)

func assertReview(t *testing.T, actual, expected *npool.Review) {
	assert.Equal(t, actual.ObjectType, expected.ObjectType)
	assert.Equal(t, actual.State, expected.State)
	assert.Equal(t, actual.Message, expected.Message)
	assert.Equal(t, actual.ObjectID, expected.ObjectID)
	assert.Equal(t, actual.Domain, expected.Domain)
}

func TestCreateReview(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	review := npool.Review{
		ObjectType: "good",
		State:      "wait",
		ObjectID:   uuid.New().String(),
		AppID:      uuid.New().String(),
		Domain:     fmt.Sprintf("cloud-hashing-goods-npool-top-%v", uuid.New().String()),
	}
	firstCreateInfo := npool.CreateReviewResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateReviewRequest{
			Info: &review,
		}).
		Post("http://localhost:50050/v1/create/review")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assertReview(t, firstCreateInfo.Info, &review)
		}
	}

	review.ID = firstCreateInfo.Info.ID
	review.State = "approved"
	review.ReviewerID = uuid.New().String()
	review.Message = "Good good good"

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateReviewRequest{
			Info: &review,
		}).
		Post("http://localhost:50050/v1/update/review")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateReviewResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertReview(t, info.Info, &review)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetReviewsByAppDomainRequest{
			AppID:  review.AppID,
			Domain: review.Domain,
		}).
		Post("http://localhost:50050/v1/get/reviews/by/app/domain")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetReviewsByAppDomainResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, len(info.Infos), 1)
		}
	}
}
