DROP TABLE IF EXISTS public.categories_items;
DROP TABLE IF EXISTS public.items;
DROP TABLE IF EXISTS public.categories;
DROP TABLE IF EXISTS public.users;

CREATE TABLE public.users
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name text NOT NULL,
    password_hash text NOT NULL,
    PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to qm_owner;

CREATE TABLE public.items
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid,
    name text,
    description text,
    image text,
    PRIMARY KEY (id),
    CONSTRAINT fk__items__client_id__cleints__id FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
)

TABLESPACE pg_default;


ALTER TABLE IF EXISTS public.items
    OWNER to qm_owner;


CREATE TABLE public.categories
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid,
    name text,
    description text,
    PRIMARY KEY (id),
    CONSTRAINT fk__categories__user_id__users__id FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.categories
    OWNER to qm_owner;

CREATE TABLE public.categories_items
(
    user_id uuid,
    category_id uuid,
    item_id uuid,
    CONSTRAINT fk__category_items__item_id__items__id FOREIGN KEY (item_id)
        REFERENCES public.items (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk__category_items__category_id__categories__id FOREIGN KEY (category_id)
        REFERENCES public.categories (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk__category_items__user_id__users__id FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);


ALTER TABLE IF EXISTS public.categories_items
    OWNER to qm_owner;

ALTER TABLE public.items ENABLE ROW LEVEL SECURITY;

ALTER TABLE public.categories ENABLE ROW LEVEL SECURITY;

ALTER TABLE public.categories_items ENABLE ROW LEVEL SECURITY;

CREATE POLICY item_owner_policy ON public.items
    USING (user_id = current_setting('app.current_user_id')::uuid);

CREATE POLICY category_owner_policy ON public.categories
    USING (user_id = current_setting('app.current_user_id')::uuid);

CREATE POLICY category_owner_policy ON public.categories_items
    USING (user_id = current_setting('app.current_user_id')::uuid);


ALTER TABLE public.users FORCE ROW LEVEL SECURITY;
ALTER TABLE public.items FORCE ROW LEVEL SECURITY;
ALTER TABLE public.categories FORCE ROW LEVEL SECURITY;
ALTER TABLE public.categories_items FORCE ROW LEVEL SECURITY;