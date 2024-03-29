package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type BookmarkUsecase interface {
	BookmarkedArticlePerPage(userId uint, pageNum int) ([]model.Article, error)
	AllBookmarkedArticle(userId uint) ([]model.Article, error)
	PostBookmark(userId uint, articleId string) (model.Bookmark, error)
}

type bookmarkUsecase struct {
	br repository.BookmarkRepository
}

func NewBookmarkUsecase(br repository.BookmarkRepository) BookmarkUsecase {
	return &bookmarkUsecase{br}
}

func (bu *bookmarkUsecase) BookmarkedArticlePerPage(userId uint, pageNum int) ([]model.Article, error) {
	bookmarkedArticles, err := bu.br.BookmarkedArticlesPerPages(userId, pageNum)
	if err != nil {
		return []model.Article{}, err
	}
	return bookmarkedArticles, nil
}

func (bu *bookmarkUsecase) AllBookmarkedArticle(userId uint) ([]model.Article, error) {
	bookmarkedArticles, err := bu.br.AllBookmarkedArticleByUserId(userId)
	if err != nil {
		return []model.Article{}, err
	}
	return bookmarkedArticles, nil
}

func (bu *bookmarkUsecase) PostBookmark(userId uint, articleId string) (model.Bookmark, error) {
	newBookmark := model.Bookmark{
		UserID:    userId,
		ArticleID: articleId,
	}
	err := bu.br.PostBookmark(&newBookmark)
	if err != nil {
		return model.Bookmark{}, err
	}
	return newBookmark, nil
}
