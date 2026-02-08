FROM migrate/migrate

COPY ./helm/migrations                          ./migrations
COPY ./docker/migrations/scripts/migration.sh   ./migration.sh


ENTRYPOINT ["./migration.sh"]
