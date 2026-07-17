
INSERT INTO `users` (`name`, `email`, `password`, `phone`, `address`, `user_type`) VALUES
('Budi Santoso', 'budi.admin@kedai.co.id', 'hashed_pass_123', '081234567890', 'Jl. Sudirman No. 10, Jakarta', 'admin'),
('Siti Aminah', 'siti.staff@kedai.co.id', 'hashed_pass_456', '085612345678', 'Jl. Melati No. 5, Surabaya', 'staff'),
('Agus Pratama', 'agus.pratama@gmail.com', 'hashed_pass_789', '081398765432', 'Jl. Merdeka No. 1, Magetan', 'customer'),
('Rina Wijaya', 'rina.wjy@yahoo.co.id', 'hashed_pass_000', '082188776655', 'Jl. Pahlawan No. 45, Madiun', 'customer'),
('Dwi Cahyono', 'dwicahyono99@gmail.com', 'hashed_pass_111', '087811223344', 'Jl. Diponegoro No. 8, Ngawi', 'customer'),
('Ayu Lestari', 'ayu.lestari_88@hotmail.com', 'hashed_pass_222', '089644556677', 'Jl. Ahmad Yani No. 12, Ponorogo', 'customer'),
('Hendra Gunawan', 'hendra.gun@gmail.com', 'hashed_pass_333', '081933445566', 'Jl. Cokroaminoto No. 3, Kediri', 'customer');


INSERT INTO `categories` (`name`) VALUES
('Coffee & Espresso'),
('Tea & Matcha'),
('Pastries & Snacks');


INSERT INTO `add_ons` (`name`, `price`) VALUES
('Extra Espresso Shot', 5000), 
('Oat Milk Upgrade', 6000), 
('Almond Milk Upgrade', 6000), 
('Vanilla Syrup', 5000), 
('Caramel Drizzle', 5000); 


INSERT INTO `payments_method` (`name`) VALUES
('Credit Card'),
('Cash'),
('QRIS / E-Wallet');


INSERT INTO `products` (`category_id`, `name`, `stock`, `price`) VALUES
(1, 'Iced Americano', 50, 20000), 
(1, 'Caramel Macchiato', 40, 28000), 
(1, 'Cold Brew Coffee', 35, 30000), 
(2, 'Iced Matcha Latte', 40, 32000), 
(2, 'Earl Grey Tea', 50, 18000), 
(3, 'Butter Croissant', 25, 15000), 
(3, 'Chocolate Chip Cookie', 30, 12000); 


INSERT INTO `product_add_ons` (`product_id`, `add_on_id`) VALUES
(1, 1), (1, 4), (1, 5), -- Americano allows Extra Shot, Vanilla, Caramel
(2, 1), (2, 2), (2, 3), -- Macchiato allows Extra Shot, Oat, Almond
(3, 4), (3, 5),         -- Cold Brew allows Vanilla, Caramel
(4, 2), (4, 3);         -- Matcha allows Oat, Almond


INSERT INTO `orders` (`user_id`, `status`, `total_amount`) VALUES

(3, 'completed', 53000), 


(4, 'preparing', 66000); 


INSERT INTO `order_items` (`order_id`, `product_id`, `add_on_id`, `quantity`, `unit_price`, `subtotal`, `note`) VALUES

(1, 4, 2, 1, 32000, 38000, 'Less ice please'), 

(1, 6, NULL, 1, 15000, 15000, 'Warm it up'), 


(2, 2, 4, 2, 28000, 66000, 'Extra sweet'); 

-- 9. Process Payments
INSERT INTO `payments` (`order_id`, `amount`, `payment_method_id`, `status`, `paid_at`) VALUES
(1, 53000, 3, 'Paid', CURRENT_TIMESTAMP), -
(2, 66000, 1, 'Pending', NULL); 