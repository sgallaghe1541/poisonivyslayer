# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
dev/templ:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
dev/server:
	air \
	--build.cmd "go build -o tmp/main" \
	--build.bin "tmp/main" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error false \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
dev/tailwind:
	./tailwindcss --input static/css/in.css --output static/css/out.css --minify --watch

# watch for any js or css change in the static/ folder, then reload the browser via templ proxy.
dev/sync_assets:
	air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "static" \
	--build.include_ext "js,css"

# start all 4 watch processes in parallel.
dev: 
	make -j2 dev/server dev/tailwind 
