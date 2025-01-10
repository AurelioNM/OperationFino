# OperationFino

### Sprint 1 - Tasks
- [x] base api - customer-service
- [x] db conn
- [x] db schema
- [x] config docker
- [x] config pyroscope
- [x] layers handler, service, resources
- [x] go Context
- [x] swagger
- Observability
    - [x] logs
    - [x] generate metrics
    - [x] config prometheus
    - [x] create basic grafana dashboards
    - [x] improve dashboards
	- [x] transform Update-Req dashboard into RequestsByStatusCode
	- [x] add reqs by status code dashboard
	- [x] add reqs rate dashboard
	- [x] add error dashboard
- Tests
    - [x] k6
	- [x] basic load tests
	- [x] test end to end apis
- Resilience
    - [x] Cache - rotas v2


### Sprint 2 - Tasks
- [ ] Customer service
    - [x] Get by name V1
    - [x] Get by name V2 with cache
    - [ ] Finish unit tests
- [ ] Product service
    - [X] CRUD operations
    - [X] Load tests
    - [X] Profiler and Metrics config
- [ ] Order service
    - [X] External requests to product and customer
    - [X] CRUD operations
    - [ ] Load tests
- Observability
    - [ ] Add metricas na camada de DB
    - [ ] Configurar grafana docker - sync entre edição dos dashs e arquivo de config
    - [ ] Melhorar metrica e dash de ReqByStatusCode
- Tests
    - [ ] Add /utils no k6 para funcoes comuns entre os testes
    - [ ] Robot
    - [ ] Test Containers
    - [ ] Vegeta
- Resilience
    - [ ] Nginx
    - [ ] Circuit Break


Mongo commands

mongosh "mongodb://order:order@of-order-mongo:27017/admin" --apiVersion 1 --username order
use order-service
db.order.find()
db.order.find({ _id: ObjectId("67806546497f5c6e81dde4ec") });


