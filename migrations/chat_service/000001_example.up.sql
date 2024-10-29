CREATE TABLE IF NOT EXISTS public.orders (
    order_id uuid NOT NULL PRIMARY KEY,
    user_id text NOT NULL
);

CREATE INDEX IF NOT EXISTS orders_user_id_idx ON public.orders (user_id);