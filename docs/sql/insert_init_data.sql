-- 插入初始的虚拟数据

use minimark;

INSERT INTO category (name) VALUES
                                ('水果'),
                                ('蔬菜'),
                                ('肉类'),
                                ('熟食'),
                                ('零食');


INSERT INTO goods (gtin, name, category_id, price, charge_type, stock) VALUES
                                                                           ('5901234123457', '苹果', 1, 3.50, 1, 150.000),
                                                                           ('5901234123458', '番茄', 2, 2.00, 1, 200.000),
                                                                           ('5901234123459', '猪肉', 3, 20.00, 2, 50.000),
                                                                           ('5901234123460', '鸡翅', 3, 10.00, 2, 80.000),
                                                                           ('5901234123461', '薯片', 5, 5.00, 2, 100.000);

INSERT INTO pay_method (name) VALUES
                                  ('现金'),
                                  ('微信支付'),
                                  ('支付宝'),
                                  ('银行卡');

INSERT INTO orders (total_paid, pay_method_id, time) VALUES
                                                        (100.00, 1, '2024-06-02 14:30:00'),
                                                        (200.00, 2, '2024-06-02 15:00:00'),
                                                        (150.00, 3, '2024-06-02 16:00:00');

INSERT INTO goods_orders (goods_id, orders_id, quantity, price) VALUES
                                                                  (1, 1, 10.000, 3.50),
                                                                  (2, 1, 5.000, 2.00),
                                                                  (4, 2, 5.000, 10.00),
                                                                  (5, 3, 10.000, 5.00);

INSERT INTO procurement (total_paid, pay_method_id, purchase_time, pay_time) VALUES
                                                                                 (1000.00, 4, '2024-06-01 09:00:00', '2024-06-01 10:00:00'),
                                                                                 (1500.00, 4, '2024-06-02 09:00:00', '2024-06-02 10:00:00');

INSERT INTO goods_procurement (goods_id, procurement_id, quantity, price) VALUES
                                                                              (1, 1, 100.000, 3.00),
                                                                              (2, 1, 150.000, 1.80),
                                                                              (3, 2, 30.000, 19.00),
                                                                              (4, 2, 60.000, 9.00);
