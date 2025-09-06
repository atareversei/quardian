-- +migrate Up
CREATE TYPE record_status AS ENUM (
    'active', -- Default state for normal operations
    'inactive', -- Not currently in use but may be reactivated
    'deleted', -- Soft-deleted (hidden from normal views)
    'pending', -- Awaiting approval/activation
    'draft', -- Incomplete configuration
    'under_review', -- Being evaluated by admins
    'suspended', -- Temporarily blocked (e.g., for investigation)
    'blacklisted', -- Permanently blocked (security reasons)
    'archived', -- Retired but kept for historical records
    'deprecated' -- Scheduled for future removal
    );

-- +migrate Down
DROP TYPE IF EXISTS record_status;