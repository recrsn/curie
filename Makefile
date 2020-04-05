e2e_server:
	$(MAKE) --directory=server build start
	$(MAKE) --directory=e2e server
	$(MAKE) --directory=server stop


e2e: e2e_server

all: e2e
