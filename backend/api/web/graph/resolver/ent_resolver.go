package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/gva/api/web/graph/generated"
	"github.com/gva/internal/ent"
)

// Chapter is the resolver for the chapter field.
func (r *comicResolver) Chapter(ctx context.Context, obj *ent.Comic) (int, error) {
	panic(fmt.Errorf("not implemented: Chapter - chapter"))
}

// Covers is the resolver for the covers field.
func (r *comicResolver) Covers(ctx context.Context, obj *ent.Comic) ([]string, error) {
	panic(fmt.Errorf("not implemented: Covers - covers"))
}

// UpCount is the resolver for the upCount field.
func (r *comicResolver) UpCount(ctx context.Context, obj *ent.Comic) (int, error) {
	panic(fmt.Errorf("not implemented: UpCount - upCount"))
}

// Chapter is the resolver for the chapter field.
func (r *comicChapterResolver) Chapter(ctx context.Context, obj *ent.ComicChapter) (int, error) {
	panic(fmt.Errorf("not implemented: Chapter - chapter"))
}

// UpCount is the resolver for the upCount field.
func (r *comicChapterResolver) UpCount(ctx context.Context, obj *ent.ComicChapter) (int, error) {
	panic(fmt.Errorf("not implemented: UpCount - upCount"))
}

// DownCount is the resolver for the downCount field.
func (r *comicChapterResolver) DownCount(ctx context.Context, obj *ent.ComicChapter) (int, error) {
	panic(fmt.Errorf("not implemented: DownCount - downCount"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Meta is the resolver for the meta field.
func (r *routeResolver) Meta(ctx context.Context, obj *ent.Route) (string, error) {
	panic(fmt.Errorf("not implemented: Meta - meta"))
}

// Chapter is the resolver for the chapter field.
func (r *comicChapterWhereInputResolver) Chapter(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: Chapter - chapter"))
}

// ChapterNeq is the resolver for the chapterNEQ field.
func (r *comicChapterWhereInputResolver) ChapterNeq(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterNeq - chapterNEQ"))
}

// ChapterIn is the resolver for the chapterIn field.
func (r *comicChapterWhereInputResolver) ChapterIn(ctx context.Context, obj *ent.ComicChapterWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: ChapterIn - chapterIn"))
}

// ChapterNotIn is the resolver for the chapterNotIn field.
func (r *comicChapterWhereInputResolver) ChapterNotIn(ctx context.Context, obj *ent.ComicChapterWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: ChapterNotIn - chapterNotIn"))
}

// ChapterGt is the resolver for the chapterGT field.
func (r *comicChapterWhereInputResolver) ChapterGt(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterGt - chapterGT"))
}

// ChapterGte is the resolver for the chapterGTE field.
func (r *comicChapterWhereInputResolver) ChapterGte(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterGte - chapterGTE"))
}

// ChapterLt is the resolver for the chapterLT field.
func (r *comicChapterWhereInputResolver) ChapterLt(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterLt - chapterLT"))
}

// ChapterLte is the resolver for the chapterLTE field.
func (r *comicChapterWhereInputResolver) ChapterLte(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterLte - chapterLTE"))
}

// UpCount is the resolver for the upCount field.
func (r *comicChapterWhereInputResolver) UpCount(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCount - upCount"))
}

// UpCountNeq is the resolver for the upCountNEQ field.
func (r *comicChapterWhereInputResolver) UpCountNeq(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountNeq - upCountNEQ"))
}

// UpCountIn is the resolver for the upCountIn field.
func (r *comicChapterWhereInputResolver) UpCountIn(ctx context.Context, obj *ent.ComicChapterWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: UpCountIn - upCountIn"))
}

// UpCountNotIn is the resolver for the upCountNotIn field.
func (r *comicChapterWhereInputResolver) UpCountNotIn(ctx context.Context, obj *ent.ComicChapterWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: UpCountNotIn - upCountNotIn"))
}

// UpCountGt is the resolver for the upCountGT field.
func (r *comicChapterWhereInputResolver) UpCountGt(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountGt - upCountGT"))
}

// UpCountGte is the resolver for the upCountGTE field.
func (r *comicChapterWhereInputResolver) UpCountGte(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountGte - upCountGTE"))
}

// UpCountLt is the resolver for the upCountLT field.
func (r *comicChapterWhereInputResolver) UpCountLt(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountLt - upCountLT"))
}

// UpCountLte is the resolver for the upCountLTE field.
func (r *comicChapterWhereInputResolver) UpCountLte(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountLte - upCountLTE"))
}

// DownCount is the resolver for the downCount field.
func (r *comicChapterWhereInputResolver) DownCount(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: DownCount - downCount"))
}

// DownCountNeq is the resolver for the downCountNEQ field.
func (r *comicChapterWhereInputResolver) DownCountNeq(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: DownCountNeq - downCountNEQ"))
}

// DownCountIn is the resolver for the downCountIn field.
func (r *comicChapterWhereInputResolver) DownCountIn(ctx context.Context, obj *ent.ComicChapterWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: DownCountIn - downCountIn"))
}

// DownCountNotIn is the resolver for the downCountNotIn field.
func (r *comicChapterWhereInputResolver) DownCountNotIn(ctx context.Context, obj *ent.ComicChapterWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: DownCountNotIn - downCountNotIn"))
}

// DownCountGt is the resolver for the downCountGT field.
func (r *comicChapterWhereInputResolver) DownCountGt(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: DownCountGt - downCountGT"))
}

// DownCountGte is the resolver for the downCountGTE field.
func (r *comicChapterWhereInputResolver) DownCountGte(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: DownCountGte - downCountGTE"))
}

// DownCountLt is the resolver for the downCountLT field.
func (r *comicChapterWhereInputResolver) DownCountLt(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: DownCountLt - downCountLT"))
}

// DownCountLte is the resolver for the downCountLTE field.
func (r *comicChapterWhereInputResolver) DownCountLte(ctx context.Context, obj *ent.ComicChapterWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: DownCountLte - downCountLTE"))
}

// Chapter is the resolver for the chapter field.
func (r *comicWhereInputResolver) Chapter(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: Chapter - chapter"))
}

// ChapterNeq is the resolver for the chapterNEQ field.
func (r *comicWhereInputResolver) ChapterNeq(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterNeq - chapterNEQ"))
}

// ChapterIn is the resolver for the chapterIn field.
func (r *comicWhereInputResolver) ChapterIn(ctx context.Context, obj *ent.ComicWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: ChapterIn - chapterIn"))
}

// ChapterNotIn is the resolver for the chapterNotIn field.
func (r *comicWhereInputResolver) ChapterNotIn(ctx context.Context, obj *ent.ComicWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: ChapterNotIn - chapterNotIn"))
}

// ChapterGt is the resolver for the chapterGT field.
func (r *comicWhereInputResolver) ChapterGt(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterGt - chapterGT"))
}

// ChapterGte is the resolver for the chapterGTE field.
func (r *comicWhereInputResolver) ChapterGte(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterGte - chapterGTE"))
}

// ChapterLt is the resolver for the chapterLT field.
func (r *comicWhereInputResolver) ChapterLt(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterLt - chapterLT"))
}

// ChapterLte is the resolver for the chapterLTE field.
func (r *comicWhereInputResolver) ChapterLte(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: ChapterLte - chapterLTE"))
}

// UpCount is the resolver for the upCount field.
func (r *comicWhereInputResolver) UpCount(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCount - upCount"))
}

// UpCountNeq is the resolver for the upCountNEQ field.
func (r *comicWhereInputResolver) UpCountNeq(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountNeq - upCountNEQ"))
}

// UpCountIn is the resolver for the upCountIn field.
func (r *comicWhereInputResolver) UpCountIn(ctx context.Context, obj *ent.ComicWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: UpCountIn - upCountIn"))
}

// UpCountNotIn is the resolver for the upCountNotIn field.
func (r *comicWhereInputResolver) UpCountNotIn(ctx context.Context, obj *ent.ComicWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: UpCountNotIn - upCountNotIn"))
}

// UpCountGt is the resolver for the upCountGT field.
func (r *comicWhereInputResolver) UpCountGt(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountGt - upCountGT"))
}

// UpCountGte is the resolver for the upCountGTE field.
func (r *comicWhereInputResolver) UpCountGte(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountGte - upCountGTE"))
}

// UpCountLt is the resolver for the upCountLT field.
func (r *comicWhereInputResolver) UpCountLt(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountLt - upCountLT"))
}

// UpCountLte is the resolver for the upCountLTE field.
func (r *comicWhereInputResolver) UpCountLte(ctx context.Context, obj *ent.ComicWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: UpCountLte - upCountLTE"))
}

// Comic returns generated.ComicResolver implementation.
func (r *Resolver) Comic() generated.ComicResolver { return &comicResolver{r} }

// ComicChapter returns generated.ComicChapterResolver implementation.
func (r *Resolver) ComicChapter() generated.ComicChapterResolver { return &comicChapterResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Route returns generated.RouteResolver implementation.
func (r *Resolver) Route() generated.RouteResolver { return &routeResolver{r} }

// ComicChapterWhereInput returns generated.ComicChapterWhereInputResolver implementation.
func (r *Resolver) ComicChapterWhereInput() generated.ComicChapterWhereInputResolver {
	return &comicChapterWhereInputResolver{r}
}

// ComicWhereInput returns generated.ComicWhereInputResolver implementation.
func (r *Resolver) ComicWhereInput() generated.ComicWhereInputResolver {
	return &comicWhereInputResolver{r}
}

type comicResolver struct{ *Resolver }
type comicChapterResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type routeResolver struct{ *Resolver }
type comicChapterWhereInputResolver struct{ *Resolver }
type comicWhereInputResolver struct{ *Resolver }