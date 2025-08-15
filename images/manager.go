package images

import (
        "fmt"
        "friday-bot/logger"
        "math/rand"
        "os"
        "path/filepath"
        "strings"
        "time"
)

// Manager handles image operations
type Manager struct {
        imagesDir string
        random    *rand.Rand
}

// New creates a new image manager
func New(imagesDir string) *Manager {
        return &Manager{
                imagesDir: imagesDir,
                random:    rand.New(rand.NewSource(time.Now().UnixNano())),
        }
}

// GetRandomImage returns a random image path from the images directory
func (m *Manager) GetRandomImage() (string, error) {
        images, err := m.listImages()
        if err != nil {
                return "", fmt.Errorf("failed to list images: %w", err)
        }

        if len(images) == 0 {
                return "", fmt.Errorf("no images found in directory: %s", m.imagesDir)
        }

        // Select random image
        selectedImage := images[m.random.Intn(len(images))]
        logger.Info("Selected image: %s", selectedImage)
        
        return selectedImage, nil
}

// listImages returns all image files in the images directory
func (m *Manager) listImages() ([]string, error) {
