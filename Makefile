
run-app:
	docker-compose -f docker-compose.yml up

clean-app:
	docker-compose -f docker-compose.yml down && docker volume rm go-profile_database_mysql