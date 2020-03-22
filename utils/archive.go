package utils

import (
	models "../models"
	config "../models/config"
	log "./log"
)

// IsItemInArchive - Checks if the item is already archived
func IsItemInArchive(upload models.Upload) bool {
	if len(RemoveArchivedItems([]models.Upload{upload})) == 0 {
		return true
	}
	return false
}

// RemoveArchivedItems - Returns items slice without archived items
func RemoveArchivedItems(uploads []models.Upload) []models.Upload {
	archiveFilePath := config.Config.ArchiveFilePath

	if archiveFilePath == "" || !CheckIfExists(archiveFilePath) {
		return uploads
	}

	removeArchivedItemsDelegate := func(archivedItem string) {
		for i, upload := range uploads {
			if upload.GetUploadID() == archivedItem {
				uploads = append(uploads[:i], uploads[i+1:]...)
			}
		}
	}

	lenBeforeRemoval := len(uploads)
	ReadFileLineByLine(archiveFilePath, removeArchivedItemsDelegate)

	removedCount := lenBeforeRemoval - len(uploads)
	if removedCount > 0 {
		log.Logf("%d items, found in archive. Skipping...\n", removedCount)
	}

	return uploads
}

// AddItemToArchive - Adds item to archived list
func AddItemToArchive(uploadID string) {
	archiveFilePath := config.Config.ArchiveFilePath

	if archiveFilePath == "" {
		return
	}

	AppendToFile(uploadID, archiveFilePath)
}
