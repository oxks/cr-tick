-- name: AddPrice :exec
INSERT INTO public.wsminimarketsstatevent ("event","time",symbol,last_price,open_price,high_price,low_price,base_volume,quote_volume)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);