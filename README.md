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
    - [ ] Paginacao 
- [ ] Definir proximo service
- [ ] Definir service worker + processamento de arquivos
- [ ] Definir service com DB nao relacional
- Observability
    - [ ] Configurar grafana docker - sync entre edição dos dashs e arquivo de config
- Tests
    - [ ] Robot
    - [ ] Test Containers
    - [ ] Vegeta
- Resilience
	- [ ] Nginx
	- [ ] Circuit Break

