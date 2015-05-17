create table  if not exists tickets (
    ticket_id text PRIMARY KEY,
    ref_id text,
    description text,
    status text,
    created_at timestamptz
)
