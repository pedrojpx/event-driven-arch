CREATE DATABASE IF NOT EXISTS `wallet`;
CREATE DATABASE IF NOT EXISTS `balances`;

USE `balances`;

CREATE TABLE `accounts` (
  `id` varchar(255) DEFAULT NULL,
  `client_id` varchar(255) DEFAULT NULL,
  `balance` int(11) DEFAULT NULL,
  `created_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES ('05f26ad8-f153-4851-b62d-bb79599b3d3e','aeab6609-114f-4598-8812-3a84861e9671',75,'2024-11-26'),('b888990e-c392-48e9-8766-e754b1041a73','d357761d-4122-446e-8736-4d3ba7aa001e',25,'2024-11-26'),('cfa72612-9f48-4642-83f1-57c715f1eccb','0b3b962e-8742-40d7-bda5-55c4b36a95ff',986,'2025-01-21'),('468562bc-d96b-4087-8fb5-255a48f3655d','475efed7-6208-4938-abaf-8a417d936952',1014,'2025-01-21');
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;



-- MySQL dump 10.13  Distrib 5.7.44, for Linux (x86_64)
--
-- Host: localhost    Database: wallet
-- ------------------------------------------------------
-- Server version       5.7.44

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `accounts`
--
USE `wallet`;

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `accounts` (
  `id` varchar(255) DEFAULT NULL,
  `client_id` varchar(255) DEFAULT NULL,
  `balance` int(11) DEFAULT NULL,
  `created_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES ('05f26ad8-f153-4851-b62d-bb79599b3d3e','aeab6609-114f-4598-8812-3a84861e9671',75,'2024-11-26'),('b888990e-c392-48e9-8766-e754b1041a73','d357761d-4122-446e-8736-4d3ba7aa001e',25,'2024-11-26'),('cfa72612-9f48-4642-83f1-57c715f1eccb','0b3b962e-8742-40d7-bda5-55c4b36a95ff',986,'2025-01-21'),('468562bc-d96b-4087-8fb5-255a48f3655d','475efed7-6208-4938-abaf-8a417d936952',1014,'2025-01-21');
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `clients`
--

DROP TABLE IF EXISTS `clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `clients` (
  `id` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clients`
--

LOCK TABLES `clients` WRITE;
/*!40000 ALTER TABLE `clients` DISABLE KEYS */;
INSERT INTO `clients` VALUES ('d357761d-4122-446e-8736-4d3ba7aa001e','John Doe','j@j.com','2024-11-26'),('aeab6609-114f-4598-8812-3a84861e9671','Jane Doe','j@j.com','2024-11-26'),('0b3b962e-8742-40d7-bda5-55c4b36a95ff','Jane Doe','j@j.com','2025-01-21'),('475efed7-6208-4938-abaf-8a417d936952','John Doe','j@j.com','2025-01-21');
/*!40000 ALTER TABLE `clients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactions` (
  `id` varchar(255) DEFAULT NULL,
  `account_id_from` varchar(255) DEFAULT NULL,
  `account_id_to` varchar(255) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `created_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('4052d7f1-796a-4a9b-b4cd-3a696d9ea36a','b888990e-c392-48e9-8766-e754b1041a73','05f26ad8-f153-4851-b62d-bb79599b3d3e',75,'2024-11-26'),('800d891e-8f4a-4929-83aa-0730662da20b','b888990e-c392-48e9-8766-e754b1041a73','05f26ad8-f153-4851-b62d-bb79599b3d3e',75,'2024-11-26'),('5451a614-6e75-46e3-832f-d28f2a5d7e4d','b888990e-c392-48e9-8766-e754b1041a73','05f26ad8-f153-4851-b62d-bb79599b3d3e',75,'2024-11-26'),('5f02af37-9e7b-47f9-b26b-df5fe3bcaa7a','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-21'),('586d0f93-6ebf-49be-be4b-947fb506cb67','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-21'),('b38e75e8-95ba-46cb-9ce1-666d44264a42','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-21'),('3652d84d-aa86-427a-a5d2-95271baf6e38','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-21'),('cdf14d58-3018-4373-a264-40765d3b1bf8','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-21'),('56514b4a-2f19-421c-a8f5-00f5be45b19e','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-21'),('9400436a-68ac-4f0b-990a-09a298b22790','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-21'),('0febaa25-8802-4be9-ab13-498ce8947edf','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-21'),('397ad22f-66d1-4799-85c6-a360f1f19af9','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-22'),('90f6bd75-830e-4a01-9e17-ee3d53a64d0b','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-29'),('1e5f1d0f-786d-4407-872a-5c6f431b7f62','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-29'),('99182466-dbc7-49b2-a29d-43c6f187c8ae','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-29'),('9b60eb61-f4a2-4bfb-ab72-55618e2932d5','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-29'),('91b20810-d6fa-4cd0-b5fc-0559496b818a','cfa72612-9f48-4642-83f1-57c715f1eccb','468562bc-d96b-4087-8fb5-255a48f3655d',1,'2025-01-29');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-01-29 23:40:51