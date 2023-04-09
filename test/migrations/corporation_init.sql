DELETE FROM corporation;
INSERT INTO corporation (id, name) VALUES ('test-corporation', 'test-corporation');

DELETE FROM menu;
INSERT INTO menu (id, menu_type, corporation_id, menu_name_en, menu_name_th) 
VALUES 
(1, 'MAIN_DISH', 'test-corporation', 'Chicken Rice', 'ไก่ย่าง'),
(2, 'MAIN_DISH', 'test-corporation', 'Pork Rice', 'หมูย่าง'),
(3, 'MAIN_DISH', 'test-corporation', 'Beef Rice', 'เนื้อย่าง'),
(4, 'MAIN_DISH', 'test-corporation', 'Fish Rice', 'ปลาทอด'),
(5, 'MAIN_DISH', 'test-corporation', 'Pork Noodle', 'หมูกระเทียม'),
(6, 'MAIN_DISH', 'test-corporation', 'Beef Noodle', 'เนื้อกระเทียม'),
(7, 'MAIN_DISH', 'test-corporation', 'Fish Noodle', 'ปลากระเทียม'),
(8, 'MAIN_DISH', 'test-corporation', 'Pork Soup', 'หมูต้ม'),
(9, 'MAIN_DISH', 'test-corporation', 'Beef Soup', 'เนื้อต้ม'),
(10, 'MAIN_DISH', 'test-corporation', 'Fish Soup', 'ปลาต้ม'),
(11, 'MAIN_DISH', 'test-corporation', 'Pork Fried Rice', 'ข้าวผัดหมู'),
(12, 'MAIN_DISH', 'test-corporation', 'Beef Fried Rice', 'ข้าวผัดเนื้อ'),
(13, 'MAIN_DISH', 'test-corporation', 'Fish Fried Rice', 'ข้าวผัดปลา'),
(14, 'MAIN_DISH', 'test-corporation', 'Pork Fried Noodle', 'ข้าวผัดหมูกระเทียม'),
(15, 'MAIN_DISH', 'test-corporation', 'Beef Fried Noodle', 'ข้าวผัดเนื้อกระเทียม'),
(16, 'MAIN_DISH', 'test-corporation', 'Fish Fried Noodle', 'ข้าวผัดปลากระเทียม')
;

DELETE FROM course;
INSERT INTO course (id, course_name, corporation_id, course_price)
VALUES 
(1, 'standard', 'test-corporation', 299),
(2, 'premium', 'test-corporation', 399),
(3, 'deluxe', 'test-corporation', 499)
;
