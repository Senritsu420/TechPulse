package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"errors"

	"gorm.io/gorm"
)

type bookmarkPersistence struct {
	db *gorm.DB
}

func NewBookmarkPersistence(db *gorm.DB) repository.BookmarkRepository {
	return &bookmarkPersistence{db}
}

func (bp *bookmarkPersistence) BookmarkedArticlesPerPages(userId uint, pageNum int) ([]model.Article, error) {
	bookmarkedArticles := []model.Article{}
	// 1ページあたりの記事数
	const pageSize = 30
	// ページ番号から、OFFSETの計算を行います。ページ番号は1から始まる
	offset := (pageNum - 1) * pageSize
	// LimitとOffsetメソッドを使ってページネーションを適用
	res := bp.db.Table("articles").
		Joins("INNER JOIN bookmarks ON articles.id = bookmarks.article_id").
		Where("bookmarks.user_id = ?", userId).
		Limit(pageSize).Offset(offset).
		Find(&bookmarkedArticles)
	if res.Error != nil {
		return []model.Article{}, res.Error
	}
	return bookmarkedArticles, nil
}

func (bp *bookmarkPersistence) AllBookmarkedArticleByUserId(userId uint) ([]model.Article, error) {
	bookmarkedArticles := []model.Article{}
	res := bp.db.Table("articles").Joins("INNER JOIN bookmarks ON articles.id = bookmarks.article_id").
		Where("bookmarks.user_id = ?", userId).
		Find(&bookmarkedArticles)
	if res.Error != nil {
		return nil, res.Error
	}
	return bookmarkedArticles, nil
}

func (bp *bookmarkPersistence) PostBookmark(bookmark *model.Bookmark) error {
	// ブックマークがすでに存在するかを確認
	existingBookmark := model.Bookmark{}
	if err := bp.db.Where("user_id = ? AND article_id = ?",
		bookmark.UserID, bookmark.ArticleID).First(&existingBookmark).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 予期しないエラーが発生した場合
			return err
		}
		// レコードが存在しない場合はブックマークを作成
		if err := bp.db.Create(bookmark).Error; err != nil {
			return err
		}
	} else {
		// レコードが存在する場合はブックマークを削除
		if err := bp.db.Delete(&existingBookmark).Error; err != nil {
			return err
		}
	}
	return nil
}
