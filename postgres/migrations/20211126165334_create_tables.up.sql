

CREATE TABLE public.wsminimarketsstatevent (
	id bigserial NOT NULL,
	"event" varchar NOT NULL,
	"time" bigint NOT NULL,
	symbol varchar NOT NULL,
	last_price varchar NOT NULL,
	open_price numeric(23, 8) NOT NULL,
	high_price numeric(25, 8) NOT NULL,
	low_price numeric(25, 8) NOT NULL,
	base_volume numeric(25, 8) NOT NULL,
	quote_volume numeric(25, 8) NOT NULL,
	CONSTRAINT wsminimarketsstatevent_pkey PRIMARY KEY (id)
);
CREATE INDEX wsminimarketsstatevent_last_price_idx ON public.wsminimarketsstatevent USING btree (last_price);
CREATE INDEX wsminimarketsstatevent_symbol_idx ON public.wsminimarketsstatevent USING btree (symbol);