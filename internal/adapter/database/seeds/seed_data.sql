-- Seed script for LoreCrafter database
-- This script populates the database with sample data for users, campaigns, and campaign_members

-- Clear existing data (if any)
DELETE FROM campaign_members;
DELETE FROM campaigns;
DELETE FROM users;

-- Seed users
-- Note: In a real environment, passwords would be properly hashed
-- These are placeholder UUIDs and dummy hashed passwords
INSERT INTO users (id, username, email, hashed_password, is_active, avatar_url, created_at, updated_at)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'dungeon_master', 'dm@example.com', '$2a$10$dummyhashfordmuser11111111111111111111111111111', true, 'https://example.com/avatars/dm.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('22222222-2222-2222-2222-222222222222', 'player_one', 'player1@example.com', '$2a$10$dummyhashforplayerone22222222222222222222222', true, 'https://example.com/avatars/player1.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('33333333-3333-3333-3333-333333333333', 'player_two', 'player2@example.com', '$2a$10$dummyhashforplayertwo33333333333333333333333', true, 'https://example.com/avatars/player2.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('44444444-4444-4444-4444-444444444444', 'player_three', 'player3@example.com', '$2a$10$dummyhashforplayerthree444444444444444444444', true, 'https://example.com/avatars/player3.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('55555555-5555-5555-5555-555555555555', 'game_master', 'gm@example.com', '$2a$10$dummyhashforgamemaster5555555555555555555555', true, 'https://example.com/avatars/gm.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Seed campaigns
INSERT INTO campaigns (id, title, setting_summary, setting, image_url, is_public, invite_code, created_by, created_at, updated_at)
VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'The Lost Mines of Phandelver', 'A D&D 5e adventure for beginners', 'In the city of Neverwinter, a dwarf named Gundren Rockseeker asked you to bring a wagon load of provisions to the rough-and-tumble settlement of Phandalin...', 'https://example.com/campaigns/lost_mines.jpg', true, 'lostmines123', '11111111-1111-1111-1111-111111111111', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Curse of Strahd', 'Gothic horror adventure in Barovia', 'Under raging storm clouds, the vampire Count Strahd von Zarovich stands silhouetted against the ancient walls of Castle Ravenloft...', 'https://example.com/campaigns/strahd.jpg', false, 'strahd456', '55555555-5555-5555-5555-555555555555', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Homebrew Fantasy Campaign', 'A custom world with unique lore', 'In the realm of Eldoria, ancient forces stir as the five kingdoms prepare for an inevitable conflict...', 'https://example.com/campaigns/homebrew.jpg', true, 'homebrew789', '11111111-1111-1111-1111-111111111111', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Seed campaign_members
INSERT INTO campaign_members (id, campaign_id, user_id, role, joined_at, last_accessed)
VALUES
    ('11111111-aaaa-aaaa-aaaa-111111111111', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '11111111-1111-1111-1111-111111111111', 'gm', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('22222222-aaaa-aaaa-aaaa-222222222222', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '22222222-2222-2222-2222-222222222222', 'player', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('33333333-aaaa-aaaa-aaaa-333333333333', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '33333333-3333-3333-3333-333333333333', 'player', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('44444444-aaaa-aaaa-aaaa-444444444444', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '44444444-4444-4444-4444-444444444444', 'player', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    
    ('55555555-bbbb-bbbb-bbbb-555555555555', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '55555555-5555-5555-5555-555555555555', 'gm', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('22222222-bbbb-bbbb-bbbb-222222222222', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '22222222-2222-2222-2222-222222222222', 'player', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('33333333-bbbb-bbbb-bbbb-333333333333', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '33333333-3333-3333-3333-333333333333', 'player', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    
    ('11111111-cccc-cccc-cccc-111111111111', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '11111111-1111-1111-1111-111111111111', 'gm', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('44444444-cccc-cccc-cccc-444444444444', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '44444444-4444-4444-4444-444444444444', 'player', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- End of seed script