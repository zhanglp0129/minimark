create database minimark;

use minimark;

create table category (
    id int primary key auto_increment comment '分类ID',
    name varchar(50) unique not null comment '分类名称'
) comment '商品分类表';

create table goods (
    id int primary key auto_increment comment '商品ID',
    gtin varchar(15) unique comment '商品条形码编号，采用EAN-13标准。如果没有条形码编号，或编号不满足EAN-13标准，则不填',
    name varchar(50) unique not null comment '商品名称',
    category_id int comment '商品分类ID',
    price decimal(10, 2) not null check ( price>0 ) comment '商品单价，必须为正数。对于散装称重的商品，单位为元/Kg；对于按件收费的商品，单位为元/个',
    charge_type tinyint not null check ( charge_type between 1 and 2) comment '收费类型：1表示散装称重；2表示按件收费',
    stock decimal(10, 3) not null check ( stock>0 ) comment '商品库存表。对于散装称重的商品，单位为Kg；对于按件收费的商品，单位为个',
    foreign key(category_id) references category(id)
) comment '商品表';

create table pay_method (
    id int primary key auto_increment comment '付款方式ID',
    name varchar(20) not null unique comment '付款方式名称'
) comment '付款方式表';

create table orders (
    id int primary key auto_increment comment '订单ID',
    total_paid decimal(10,2) not null check ( total_paid>=0 ) comment '合计实付金额，优惠后',
    pay_method_id int not null comment '付款方式ID',
    pay_time datetime not null comment '付款时间',
    foreign key(pay_method_id) references pay_method(id)
) comment '订单表';

create table goods_orders (
    goods_id int comment '商品ID',
    orders_id int comment '订单ID',
    quantity decimal(10, 3) not null check ( quantity>0 ) comment '购买数量：对于散装称重的商品，单位为Kg；对于按件收费的商品，单位为个',
    price decimal(10,2) not null check ( price>0 ) comment '商品单价。考虑到以后可能会涨价和降价，这里必须记录',
    primary key(goods_id, orders_id),
    foreign key(goods_id) references goods(id),
    foreign key(orders_id) references orders(id)
) comment '订单与商品的关系表';

create table procurement (
    id int primary key auto_increment comment '进货ID',
    total_paid decimal(10,2) not null check ( total_paid>=0 ) comment '合计支付金额',
    pay_method_id int not null comment '付款方式ID',
    purchase_time datetime not null comment '进货时间',
    pay_time datetime not null comment '打款时间',
    foreign key(pay_method_id) references pay_method(id)
) comment '进货表';

create table goods_procurement (
    goods_id int comment '商品ID',
    procurement_id int comment '进货ID',
    quantity decimal(10, 3) not null check ( quantity>0 ) comment '进货数量：对于散装称重的商品，单位为Kg；对于按件收费的商品，单位为个',
    price decimal(10,2) not null check ( price>0 ) comment '进货单价',
    primary key(goods_id, procurement_id),
    foreign key(goods_id) references goods(id),
    foreign key(procurement_id) references procurement(id)
) comment '商品与进货的关系表';