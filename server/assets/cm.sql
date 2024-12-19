SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS `c_m` DEFAULT  CHARACTER SET  utf8mb4;

USE c_m;



DROP TABLE IF EXISTS `computer`;
CREATE TABLE computer (
                          id INT AUTO_INCREMENT PRIMARY KEY COMMENT '商品唯一标识',
                          name VARCHAR(255) NOT NULL COMMENT '商品名称',
                          image      VARCHAR(255) NOT NULL  COMMENT '商品图片',
                          price DECIMAL(10, 2) NOT NULL COMMENT '商品价格',
                          description  VARCHAR(255) NOT NULL  COMMENT '电脑介绍',
                          created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          updated_at DATETIME DEFAULT NULL COMMENT '更新时间',
                        
) COMMENT='商品表';

DROP TABLE IF EXISTS `product_detail`;
CREATE TABLE product_detail (
                                id INT AUTO_INCREMENT PRIMARY KEY COMMENT '商品详情唯一标识',
                                product_id INT NOT NULL UNIQUE COMMENT '商品ID',
                                processor VARCHAR(255) NOT NULL COMMENT '处理器信息',
                                stock INT NOT NULL COMMENT '库存数量',
                                memory VARCHAR(255) NOT NULL COMMENT '内存信息',
                                hard_disk VARCHAR(255) NOT NULL COMMENT '硬盘信息',
                                graphics_card VARCHAR(255) NOT NULL COMMENT '显卡信息',
                                created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                updated_at DATETIME DEFAULT NULL COMMENT '更新时间',
                                FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
                                INDEX idx_product_id (product_id)
) COMMENT='商品详情表';

DROP TABLE IF EXISTS `product_image`;
CREATE TABLE product_image (
                               id INT AUTO_INCREMENT PRIMARY KEY COMMENT '图片唯一标识',
                               product_id INT NOT NULL COMMENT '商品ID',
                               image_url VARCHAR(255) NOT NULL COMMENT '图片URL',
                               description VARCHAR(255) DEFAULT NULL COMMENT '图片描述',
                               created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               updated_at DATETIME DEFAULT NULL COMMENT '更新时间',
                               FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
                               INDEX idx_product_id (product_id)
) COMMENT='商品图片表';

CREATE TABLE carousel (
                          id INT AUTO_INCREMENT PRIMARY KEY COMMENT '轮播图唯一标识',
                          image_url VARCHAR(255) NOT NULL COMMENT '图片URL',
                          link_url VARCHAR(255) DEFAULT NULL COMMENT '跳转链接',
                          is_active TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用（1：启用，0：禁用）',
                          `order` INT NOT NULL DEFAULT 0 COMMENT '显示顺序',
                          created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          updated_at DATETIME DEFAULT NULL COMMENT '更新时间'
) COMMENT='轮播图表';

DROP TABLE IF EXISTS `user`;
CREATE TABLE user (
                      id INT AUTO_INCREMENT PRIMARY KEY COMMENT '用户唯一标识',
                      username VARCHAR(255) NOT NULL UNIQUE COMMENT '用户名',
                      password VARCHAR(255) NOT NULL COMMENT '密码（加密存储）',
                      email VARCHAR(255) NOT NULL UNIQUE COMMENT '邮箱',
                      introduction VARCHAR(255) DEFAULT NULL COMMENT '简介',
                      avatar VARCHAR(255) DEFAULT NULL COMMENT '头像',
                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                      updated_at DATETIME DEFAULT NULL COMMENT '更新时间',
                      INDEX idx_username (username),
                      INDEX idx_email (email)
) COMMENT='用户表';

INSERT INTO computer (name, image, price, description, created_at, updated_at)
VALUES
('Dell XPS 13', 'https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/Carousel2.jpg', 999.99, '13-inch ultrabook with a high-resolution display', NOW(), NOW()),
('Apple MacBook Pro 16"', 'https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/Carousel2.jpg', 2499.99, '16-inch laptop with an M1 Pro chip', NOW(), NOW()),
('HP Spectre x360', 'https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/HP.jpg', 1349.99, '2-in-1 convertible laptop with a touchscreen', NOW(), NOW()),
('Lenovo ThinkPad X1 Carbon', 'https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/Carousel1.jpg', 1899.99, 'Ultra-lightweight business laptop', NOW(), NOW());

-- 插入 product_detail 表的数据
INSERT INTO product_detail (product_id, processor, stock, memory, hard_disk, graphics_card, created_at, updated_at)
VALUES
(1, 'Intel Core i7', 100, '16GB RAM', '512GB SSD', 'Intel Iris Xe', NOW(), NOW()),
(2, 'Apple M1 Pro', 50, '32GB RAM', '1TB SSD', 'Integrated GPU', NOW(), NOW()),
(3, 'Intel Core i5', 150, '8GB RAM', '256GB SSD', 'Intel UHD Graphics', NOW(), NOW()),
(4, 'Intel Core i7', 75, '16GB RAM', '512GB SSD', 'Intel Iris Xe', NOW(), NOW());

-- 插入 product_image 表的数据
INSERT INTO product_image (product_id, image_url, description, created_at, updated_at)
VALUES
(1, 'https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/DELL_image.jpg', 'Front view of Dell XPS 13', NOW(), NOW()),
(2, 'https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/Apple_image.jpg', 'Apple MacBook Pro 16" from the back', NOW(), NOW()),
(3, 'https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/hp_image.jpg', 'HP Spectre x360 laptop mode', NOW(), NOW()),
(4, 'https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/levno_image.jpg', 'Close-up of ThinkPad X1 keyboard', NOW(), NOW());

-- 插入 carousel 表的数据
INSERT INTO carousel (image_url, link_url, is_active, `order`, created_at, updated_at)
VALUES
('https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/Carousel2.jpg', 'localhost:5173/product/1', 1, 1, NOW(), NOW()),
('https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/Carousel2.jpg', 'localhost:5173/product/2', 1, 2, NOW(), NOW()),
('https://gvbresource.oss-cn-hongkong.aliyuncs.com/computer_image/HP.jpg', 'localhost:5173/product/3', 1, 3, NOW(), NOW());

-- 插入 user 表的数据
INSERT INTO user (username, password, email, introduction, avatar, created_at, updated_at)
VALUES
('john_doe', '$2a$10$FmU4jxwDlibSL9pdt.AsuODkbB4gLp3IyyXeoMmW/XALtT/HdwTsi', 'test@qq.com', 'Tech enthusiast', 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png', NOW(), NOW());

SET FOREIGN_KEY_CHECKS  = 1;