package review

import (
	"context"
	"time"

	"github.com/NpoolPlatform/review-service/message/npool"

	"github.com/NpoolPlatform/review-service/pkg/db"
	"github.com/NpoolPlatform/review-service/pkg/db/ent"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateReview(info *npool.Review) error {
	if _, err := uuid.Parse(info.GetObjectID()); err != nil {
		return xerrors.Errorf("invalid object id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if info.GetDomain() == "" {
		return xerrors.Errorf("invalid domain")
	}
	return nil
}

func dbRowToReview(row *ent.Review) *npool.Review {
	return &npool.Review{
		ID:         row.ID.String(),
		ObjectType: row.ObjectType,
		AppID:      row.AppID.String(),
		ReviewerID: row.ReviewerID.String(),
		State:      string(row.State),
		Message:    row.Message,
		ObjectID:   row.ObjectID.String(),
		Domain:     row.Domain,
	}
}

func Create(ctx context.Context, in *npool.CreateReviewRequest) (*npool.CreateReviewResponse, error) {
	if err := validateReview(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid patameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Review.
		Create().
		SetObjectType(in.GetInfo().GetObjectType()).
		SetState("wait").
		SetMessage("").
		SetReviewerID(uuid.UUID{}).
		SetObjectID(uuid.MustParse(in.GetInfo().GetObjectID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetDomain(in.GetInfo().GetDomain()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create review: %v", err)
	}

	return &npool.CreateReviewResponse{
		Info: dbRowToReview(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateReviewRequest) (*npool.UpdateReviewResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	reviewerID, err := uuid.Parse(in.GetInfo().GetReviewerID())
	if err != nil {
		return nil, xerrors.Errorf("invalid reviewer id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Review.
		UpdateOneID(id).
		SetState(review.State(in.GetInfo().GetState())).
		SetMessage(in.GetInfo().GetMessage()).
		SetReviewerID(reviewerID).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update review: %v", err)
	}

	return &npool.UpdateReviewResponse{
		Info: dbRowToReview(info),
	}, nil
}

func GetByAppDomain(ctx context.Context, in *npool.GetReviewsByAppDomainRequest) (*npool.GetReviewsByAppDomainResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Review.
		Query().
		Where(
			review.And(
				review.Domain(in.GetDomain()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query review: %v", err)
	}

	reviews := []*npool.Review{}
	for _, info := range infos {
		reviews = append(reviews, dbRowToReview(info))
	}

	return &npool.GetReviewsByAppDomainResponse{
		Infos: reviews,
	}, nil
}
