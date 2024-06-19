create table categories
(
    id   bigint auto_increment primary key,
    name varchar(20) not null,
    constraint uni_categories_name unique (name)
);

create table goods
(
    id          bigint auto_increment primary key,
    gtin        varchar(15)      null,
    name        varchar(50)      not null,
    category_id bigint           null,
    price       decimal(10, 2)   not null,
    charge_type tinyint unsigned not null,
    stock       decimal(10, 2)   not null,
    constraint uni_goods_name unique (name),
    constraint fk_goods_category foreign key (category_id) references categories (id),
    constraint chk_goods_charge_type check (`charge_type` <= 1),
    constraint chk_goods_price check (`price` > 0),
    constraint chk_goods_stock check (`stock` > 0)
);

create table pay_methods
(
    id   bigint auto_increment primary key,
    name varchar(20) not null,
    constraint uni_pay_methods_name unique (name)
);

create table orders
(
    id            bigint auto_increment primary key,
    total_paid    decimal(10, 2) not null,
    pay_method_id bigint         null,
    pay_time      datetime(3)    not null,
    constraint fk_orders_pay_method foreign key (pay_method_id) references pay_methods (id),
    constraint chk_orders_total_paid check (`total_paid` > 0)
);

create table order_goods
(
    goods_id bigint         not null,
    order_id bigint         not null,
    quantity decimal(10, 3) not null,
    price    decimal(10, 2) not null,
    primary key (goods_id, order_id),
    constraint fk_order_goods_goods foreign key (goods_id) references goods (id),
    constraint fk_order_goods_order foreign key (order_id) references orders (id),
    constraint chk_order_goods_price check (`price` > 0),
    constraint chk_order_goods_quantity check (`quantity` > 0)
);

create table procurements
(
    id            bigint auto_increment primary key,
    total_paid    decimal(10, 2) not null,
    pay_method_id bigint         null,
    purchase_time datetime(3)    not null,
    pay_time      datetime(3)    not null,
    constraint fk_procurements_pay_method foreign key (pay_method_id) references pay_methods (id),
    constraint chk_procurements_total_paid check (`total_paid` > 0)
);

create table procurement_goods
(
    goods_id       bigint         not null,
    procurement_id bigint         not null,
    quantity       decimal(10, 3) not null,
    price          decimal(10, 2) not null,
    primary key (goods_id, procurement_id),
    constraint fk_procurement_goods_goods foreign key (goods_id) references goods (id),
    constraint fk_procurement_goods_procurement foreign key (procurement_id) references procurements (id),
    constraint chk_procurement_goods_price check (`price` > 0),
    constraint chk_procurement_goods_quantity check (`quantity` > 0)
);

