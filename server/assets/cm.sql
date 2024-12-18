SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS `c_m` DEFAULT  CHARACTER SET  utf8mb4;

USE c_m;

DROP TABLE IF EXISTS `brand`;
CREATE TABLE brand (
                       id INT AUTO_INCREMENT PRIMARY KEY COMMENT '品牌唯一标识',
                       name VARCHAR(255) NOT NULL UNIQUE COMMENT '品牌名称',
                       created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       updated_at DATETIME DEFAULT NULL COMMENT '更新时间'
) COMMENT='品牌表';

DROP TABLE IF EXISTS `category`;
CREATE TABLE category (
                          id INT AUTO_INCREMENT PRIMARY KEY COMMENT '分类唯一标识',
                          name VARCHAR(255) NOT NULL UNIQUE COMMENT '分类名称',
                          created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          updated_at DATETIME DEFAULT NULL COMMENT '更新时间'
) COMMENT='分类表';

DROP TABLE IF EXISTS `computer`;
CREATE TABLE computer (
                          id INT AUTO_INCREMENT PRIMARY KEY COMMENT '商品唯一标识',
                          name VARCHAR(255) NOT NULL COMMENT '商品名称',
                          image      VARCHAR(255) NOT NULL  COMMENT '商品图片',
                          price DECIMAL(10, 2) NOT NULL COMMENT '商品价格',
                          description  VARCHAR(255) NOT NULL  COMMENT '电脑介绍',
                          brand_id INT NOT NULL COMMENT '品牌ID',
                          category_id INT NOT NULL COMMENT '分类ID',
                          created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          updated_at DATETIME DEFAULT NULL COMMENT '更新时间',
                          FOREIGN KEY (brand_id) REFERENCES brand(id) ON DELETE CASCADE ON UPDATE CASCADE,
                          FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE,
                          INDEX idx_brand_id (brand_id),
                          INDEX idx_category_id (category_id)
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

DROP TABLE IF EXISTS `role`;
CREATE TABLE role (
                      id INT AUTO_INCREMENT PRIMARY KEY COMMENT '角色唯一标识',
                      name VARCHAR(255) NOT NULL UNIQUE COMMENT '角色名称',
                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                      updated_at DATETIME DEFAULT NULL COMMENT '更新时间',
                      INDEX idx_name (name)
) COMMENT='角色表';

DROP TABLE IF EXISTS `permission`;
CREATE TABLE permission (
                            id INT AUTO_INCREMENT PRIMARY KEY COMMENT '权限唯一标识',
                            name VARCHAR(255) NOT NULL UNIQUE COMMENT '权限名称',
                            url VARCHAR(255) NOT NULL COMMENT '资源名称（如 product, order）',
                            method VARCHAR(255) NOT NULL COMMENT '操作名称（如 GET,POST）',
                            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            updated_at DATETIME DEFAULT NULL COMMENT '更新时间',
                            INDEX idx_name (name)
) COMMENT='权限表';

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE user_role (
                           id INT AUTO_INCREMENT PRIMARY KEY COMMENT '关联唯一标识',
                           user_id INT NOT NULL COMMENT '用户ID',
                           role_id INT NOT NULL COMMENT '角色ID',
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE,
                           FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE ON UPDATE CASCADE,
                           INDEX idx_user_id (user_id),
                           INDEX idx_role_id (role_id)
) COMMENT='用户角色关联表';

DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE role_permission (
                                 id INT AUTO_INCREMENT PRIMARY KEY COMMENT '关联唯一标识',
                                 role_id INT NOT NULL COMMENT '角色ID',
                                 permission_id INT NOT NULL COMMENT '权限ID',
                                 created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE ON UPDATE CASCADE,
                                 FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE ON UPDATE CASCADE,
                                 INDEX idx_role_id (role_id),
                                 INDEX idx_permission_id (permission_id)
) COMMENT='角色权限关联表';


SET FOREIGN_KEY_CHECKS  = 1;