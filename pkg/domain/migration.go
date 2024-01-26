package domain

func (UserInterests) Migration() string {
	return `
        CREATE UNIQUE INDEX idx_user_interest ON user_interests (user_id, interest_id);
    `
}

func (Images) Migration() string {
	return `
        CREATE UNIQUE INDEX idx_images ON images (user_id, image);
    `
}