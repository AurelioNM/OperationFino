drop table if exists cliente cascade; 
drop table if exists telefone cascade;
drop table if exists produto cascade;
drop table if exists pedido cascade;
drop table if exists pedido_produto cascade;
drop table if exists endereco cascade;



create table if not exists cliente (
	id_cliente int primary KEY,
	nm_cliente varchar(200),
	dt_nascimento date
);


create table if not exists telefone(
	id_telefone int primary KEY,
	id_cliente int ,
	telefone varchar(255),
	tp_telefone varchar(50),
	CONSTRAINT id_cliente_fk
	   FOREIGN KEY(id_cliente) 
	      REFERENCES cliente(id_cliente)
);

create table if not exists produto(
	id_produto int primary KEY,
	sku_produto varchar(10),
	nm_produto varchar(200),
	valor_produto decimal,
	medida_produto varchar(200),
	dt_atualizacao date,
	dt_inclucao date
);


create table if not exists pedido(
	id_pedido int primary KEY,
	id_cliente int,
	valor_pedido decimal,
	dt_pedido date,
	CONSTRAINT id_cliente_fk
	   FOREIGN KEY(id_cliente) 
	      REFERENCES cliente(id_cliente)
);

create table if not exists pedido_produto (
	id_pedido int ,
	id_produto int,
	valor_produto decimal,
	dt_pedido date,
	CONSTRAINT id_pedido_fk
	   FOREIGN KEY(id_pedido) 
	      REFERENCES pedido(id_pedido),
	CONSTRAINT id_produto_fk
	   FOREIGN KEY(id_produto) 
	      REFERENCES produto(id_produto)
);


create table if not exists endereco (
	id_endereco int primary key,
	id_cliente int,
	logradouro varchar(200),
	tp_logradouro varchar(200),
	nu_end varchar(10),
	tp_endereco int
);


insert into cliente values (1,'Felipe Cabral',to_date('1988-01-01','YYYY-MM-DD'));
insert into cliente values (2,'Arnaldo Cesar',to_date('1930-01-01','YYYY-MM-DD'));
insert into cliente values (3,'Gioconda Brasil',to_date('1940-01-01','YYYY-MM-DD'));
insert into cliente values (4,'Gustavo Lima',to_date('1960-01-01','YYYY-MM-DD'));
insert into cliente values (5,'Felipe Dylon',to_date('1980-01-01','YYYY-MM-DD'));


insert into telefone values(1,1,'8123991924','CELULAR');
insert into telefone values(2,1,'8945901233','FIXO');
insert into telefone values(3,1,'9939385829','TRABALHO');
insert into telefone values(4,2,'2988854930','CELULAR');
insert into telefone values(5,2,'2342344432','FIXO');
insert into telefone values(6,3,'4422234234','FIXO');
insert into telefone values(7,3,'5425224323','FIXO');
insert into telefone values(8,4,'8838384475','CELULAR');
insert into telefone values(9,4,'2342223423','FIXO');
insert into telefone values(10,5,'999999993234','CELULAR');
insert into telefone values(11,5,'823989238833','FIXO');
insert into telefone values(12,5,'444909595959','CELULAR');
insert into telefone values(13,2,'494949949249','FIXO');
insert into telefone values(14,1,'278273468782','CELULAR');
insert into telefone values(15,1,'993939393939','CELULAR');

insert into PRODUTO values (515276, 'SK515276','ABRIDOR DE GARRAFA QUANTA QTAGM01 MAGNETICO SILVER',3.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (770156, 'SK770156','ABRIDOR DE VINHO ELETRICO ONIDA ON-025 INOX/REC/2V',15.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (823661, 'SK823661','ABRIDOR DE VINHO ELETRICO ONIDA ON-026 KIT/RECARREGAVEL INOX 2V',16.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (785181, 'SK785181','ABRIDOR DE VINHO ELETRICO QUANTA QTAVE89N PRATA',12.90,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (868341, 'SK868341','ABRIDOR DE VINHO ELETRICO QUANTA QTAVE91 500MAH PRETO',12.90,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (824064, 'SK824064','ABRIDOR VINHO SATELLITE A-WP010 ELETRICO PRETO',13.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (754484, 'SK754484','ACESS SPORT CAM QUANTA QTSCA10 1TORNILLO+BRAZO PRETO',1.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (847117, 'SK847117','ACOPLE DE BATERIA INGCO HBTCP6001 3METRO 10MM P/ CARRO',15.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (330114, 'SK330114','ADAP CINEBOX WIFI+USB+UNIVERSAL',1.90,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (744065, 'SK744065','ADAP DE AUDIO QUANTA QTRABT10 3.5MM/USB/AUX/BT',6.40,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (795838, 'SK795838','ADAP DE REDE SATELITE LAN 01 P/ LAN RJ45',4.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (802857, 'SK802857','ADAP HYE ENCHUFE HYEC35 CARGA RAPIDA USB-C/3.0A BRANCO NOVO*',8.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (803199, 'SK803199','ADAP HYE ENCHUFE HYEC43 CARGA RAPIDA USB-C 30W BRANCO NOVO*',9.75,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (803205, 'SK803205','ADAP HYE ENCHUFE HYEC44 CARGA RAPIDA USB/USB-C 30W BRANCO NOVO*',10.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (802833, 'SK802833','ADAP HYE ENCHUFE HYEC7 CARGA RAPIDA USB/3.0A BRANCO NOVO*',7.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (802840, 'SK802840','ADAP HYE ENCHUFE HYEC9 CARGA RAPIDA USB/USB-C/3.0A BRANCO NOVO*',8.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (761406, 'SK761406','ADAP MAGNAVOX BY PHILIPS ENCHUFE MAC6619-MO CARGA RAPIDA USB/3.0A',6.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (583817, 'SK583817','ADAP QUANTA HDMI QTHDV34 P/VGA PRETO.',6.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (810005, 'SK810005','ADAP TOMADA AIWA USB AW-CSWC4P 4 SAIDA USB BIVOLT BRANCO',9.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (712699, 'SK712699','ADAP TOMADA AIWA USB AW-P2232QW 3 SAIDA USB BIVOLT BRANCO',10.90,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (696371, 'SK696371','ADAP TOMADA XIAOMI MI SMART PLUG ZNCZ05CM 220-240V BRANCO',9.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (749978, 'SK749978','ADAPTADOR BT/WIFI ECOPOWER P/PC EP-W001 1300MBPS USB UNIVERSAL',9.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (576932, 'SK576932','ADAPTADOR DE AUDIO MICRO BOYA BY-MP4',19.90,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (842792, 'SK842792','ADAPTADOR SATELITE AL-HF1431 USB-C P/ HDMI FEMEA 14CM 4K',10.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (476201, 'SK476201','ADAPTADOR WIFI PERA PR-801 USB COM ANTENA UNIVERSAL',4.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (280181, 'SK280181','ADAPTADOR WIFI PERA PR-802 USB COM ANTENA UNIVERSAL',4.25,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (681742, 'SK681742','ADAPTADOR WIFI PERA PR-806 USB COM ANTENA UNIVERSAL',4.75,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (749435, 'SK749435','ADAPTADOR WIFI QUANTA QTA802 RECEPTOR 150MBPS USB C/ANTE UNIVERSAL',4.55,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (828390, 'SK828390','ADAPTADOR WIFI TOTOLINK UNIVERSAL 150MBPS RECEPTORES NOVO*',4.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (863964, 'SK863964','ALARME CARRO HINOR HA KEYLESS UNIVERSAL',42.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (863957, 'SK863957','ALARME CARRO HINOR HA19 UNIVERSAL C/ CHAVE CANIVETE',36.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (826396, 'SK826396','ALARME CARRO POSITRON CYBER FX-360 UNIVERSAL PRETO',53.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (393577, 'SK393577','ALARME CARRO QUANTA QTAC45 COM SENSOR',15.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (782425, 'SK782425','ALICATE CRIMPADOR E DESENCAPADOR NET WORK TOOL',8.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (847087, 'SK847087','ALICATE MULTIFUNCION INGCO HCP28168 6/160MM CR-V',3.50,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (847070, 'SK847070','ALICATE MULTIFUNCION INGCO HPP28251 10/250MM',5.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (847100, 'SK847100','ALICATE MULTIFUNCION INGCO HPP28258 10/250MM CR-V',7.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (823869, 'SK823869','ALIMENTADOR XIAOMI SMART PET FOOD FEEDER XWPF01MG-EU BHR6143EU BRANCO',82.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (823876, 'SK823876','ALIMENTADOR XIAOMI SMART PET FOOD FOUNTAIN XWWF01MG-EU BHR6161EU BRANC',58.75,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (810067, 'SK810067','ALISADOR AIWA AW-TP0125 CERAMIC 220V 50/60HZ PRETO',11.90,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (810074, 'SK810074','ALISADOR AIWA AW-TP5538 ION TECH 220V 50/60HZ PRETO',18.90,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (868396, 'SK868396','ALISADOR BABYLISS + SECADOR BNTBGPP43UZ KIT LIMITED EDITION NANO GOLD',120.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (868402, 'SK868402','ALISADOR BABYLISS + SECADOR BNTPP45UZ KIT 3PCS 2V CELESTE',120.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (815994, 'SK815994','ALISADOR BABYLISS GXT GRAPHITE TITANIUM PRETO 220V BGT183UC 6422  1-1',88.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (816014, 'SK816014','ALISADOR BABYLISS NANO BNT4083TUZ 38MM AZUL BIVOLT 6731',82.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (868419, 'SK868419','ALISADOR BABYLISS PRO BNT4094TUX IONIC AZUL BIVOLT',85.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));
insert into PRODUTO values (855204, 'SK855204','ALISADOR BABYLISS PRO NANO TITANIUM OPTIMA 3300 BIVOLT',120.00,'UNIDADE', TO_DATE('2023-01-01','YYYY-MM-DD'),TO_DATE('2023-01-01','YYYY-MM-DD'));



insert into pedido values (1,1,250.00, to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido values (2,1,350.00, to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido values (3,2,540.00, to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido values (4,3,550.00, to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido values (5,4,450.00, to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido values (6,4,5660.00, to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido values (7,5,570.00, to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido values (8,5,530.00, to_date('2023-01-01','YYYY-MM-DD'));

insert into pedido_produto VALUES (1,280181,12.66,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,681742,99.66,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,749435,12.22,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,828390,33.55,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,863964,3444.33,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,863957,43223.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,826396,3234.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,393577,23.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,782425,55.42,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,847087,234.22,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,847070,4323.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,847100,33.33,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,823869,12.21,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,823876,34.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,810067,55.66,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (1,855204,88.66,to_date('2023-01-01','YYYY-MM-DD'));

insert into pedido_produto VALUES (2,280181,12.66,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (2,681742,99.66,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (2,749435,12.22,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (2,828390,33.55,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (2,863964,3444.33,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (2,863957,43223.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (2,826396,3234.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (3,393577,23.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (3,782425,55.42,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (3,847087,234.22,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (3,847070,4323.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (3,847100,33.33,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (4,823869,12.21,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (4,823876,34.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (4,810067,55.66,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (4,855204,88.66,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,863957,43223.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,826396,3234.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,393577,23.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,782425,55.42,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,847087,234.22,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,847070,4323.44,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,847100,33.33,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,823869,12.21,to_date('2023-01-01','YYYY-MM-DD'));
insert into pedido_produto VALUES (5,823876,34.44,to_date('2023-01-01','YYYY-MM-DD'));




insert into endereco values (1,1,'Nilo Pecanha','RUA','10',1);
insert into endereco values (2,1,'Jabaquara','Avenida','22',2);
insert into endereco values (3,2,'Bom Senhor Macedo','Travessa','S/N',1);
insert into endereco values (4,3,'Google da Silva da Cunha','Avenia','55555',2);
insert into endereco values (5,4,'Maria Joaquina','Rua','2222',1);
insert into endereco values (6,5,'Joao Albuquerque','Avenida','2',2);
insert into endereco values (7,4,'Joao Maria e Jose','Avenida','223',1);
insert into endereco values (8,5,'Ajaio da Silva Modeira Pedeira','Travessa','222',1);
insert into endereco values (9,1,'Rogai Por Nois','Beco','12332',2);
insert into endereco values (10,2,'Olha nos aqui novamente','Beco','2222',1);
