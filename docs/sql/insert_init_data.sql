-- 插入初始的虚拟数据

use minimark;

INSERT INTO categories (name) VALUES
                                ('水果'),
                                ('蔬菜'),
                                ('肉类'),
                                ('熟食'),
                                ('零食');


INSERT INTO goods (gtin, name, category_id, price, charge_type, stock) VALUES
                                                                           ('5901234123457', '苹果', 1, 3.50, 0, 150.000),
                                                                           ('5901234123458', '番茄', 2, 2.00, 0, 200.000),
                                                                           ('5901234123459', '猪肉', 3, 20.00, 1, 50.000),
                                                                           ('5901234123460', '鸡翅', 3, 10.00, 1, 80.000),
                                                                           ('5901234123461', '薯片', 5, 5.00, 1, 100.000);

INSERT INTO pay_methods (name) VALUES
                                  ('现金'),
                                  ('微信支付'),
                                  ('支付宝'),
                                  ('银行卡');
