.PHONY: runall
runall:
	cd deployments && docker compose up --build -d
