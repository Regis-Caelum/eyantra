CREATE DATABASE  IF NOT EXISTS `test_schema` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `test_schema`;
-- MySQL dump 10.13  Distrib 8.0.25, for Linux (x86_64)
--
-- Host: localhost    Database: test_schema
-- ------------------------------------------------------
-- Server version	8.0.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `data`
--

DROP TABLE IF EXISTS `data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `data` (
  `district` varchar(255) NOT NULL,
  `pincode` varchar(255) NOT NULL,
  `namess` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data`
--

LOCK TABLES `data` WRITE;
/*!40000 ALTER TABLE `data` DISABLE KEYS */;
INSERT INTO `data` VALUES ('Nagpur','440018','IGGMC'),('Nagpur','440030','Alexis Hospital'),('Nagpur','440013','Metro Hospital'),('District','Pincode','Hospital Name'),('Nagpur','440009','GMC'),('Nagpur','440010','Indira Gandhi Rugnalaya Gandhi Nagar'),('Nagpur','440001','Ayush Hopital (NMC Sadar)'),('Nagpur','440019','Lata Mangeshkar Hospital Hingna (S.G.)'),('Nagpur','440012','Lata Mangeshkar Hospital Sitabuldi (S.G.)'),('Nagpur','441110','Shalinitai Meghe Hospital (S.G.)'),('Nagpur','440006','K.T. Nagar NMC Hospital'),('Nagpur','440015','Orange City Hospital'),('Nagpur','440025','Grace Ortho Hospital'),('Nagpur','440022','Sanjeevani Hospital (Laxminagar)'),('Nagpur','440070','Shraman Hospital'),('Nagpur','440027','Shri Saikripa Hospital'),('Nagpur','440024','Star City Hospital'),('Nagpur','440034','Borkar Multispeciality Hospital'),('Nagpur','440017','Shravan Hospital'),('Nagpur','440002','Sawarkar Multispeciality Hospital'),('Nagpur','440008','Central Avenue Critical Care Hospital'),('Nagpur','440035','Tarangan Hospital'),('Nagpur','440014','Rugwani Child Care Center & Hospital'),('Nagpur','440023','Gajbe Critical Care & Multispeciality'),('Nagpur','440037','New Life Hospital (Unit of ChhajedHealthcare Pvt. Ltd.)'),('Nagpur','440016','Akruti Hospital'),('Nagpur','441108','National Cancer Institute Jamtha'),('Nagpur','441001','Military Hospital, Kamptee'),('Nagpur','441002','Life Line Hospital, Kamptee'),('Nagpur','441203','Arch Angel Hspital, Umred'),('Nagpur','441122','Choudhari Multispeciality Hospital, Butibori'),('Nagpur','441107','Shreyas Patil Multispeciality Hospital, Saoner'),('Nagpur','441106','Yogiraj Swami Sitaramdasji Maharaj Hospital, Ramtek'),('Nagpur','441104','BAGHE MULTISPECIALITY HOSPITAL'),('Nagpur','441113','Shourya Hospital'),('Nagpur','441302','Dhanwantari Katol'),('Nagpur','441501','Pande Childcare, Kalmeshwar'),('Nagpur','441404','Jawaharlal Neharu Central Hospital, Kandri Kanhan Parseoni'),('Nagpur','441111','MCH Koradi'),('Nagpur','441201','RH Bhiwapur'),('Nagpur','440001','Railway Hospital'),('Nagpur','440015','Kalpavruksha Hospital'),('Nagpur','440015','Swasthyam Superspeciality Hospital'),('Nagpur','440022','Viveka Superspeciality Hospital'),('Nagpur','440010','Ganga Care Hospital'),('Nagpur','440010','Zenith Hospital'),('Nagpur','440010','Central Neurological & Medical Institute'),('Nagpur','440010','G.B. Multicare Hospital'),('Nagpur','440010','Dande Hospital'),('Nagpur','440001','Mure Memorial Hospital'),('Nagpur','440012','Kalamkar Multispeciality Hospital'),('Nagpur','440009','Center Point Hospital'),('Nagpur','440012','Green City Nursing Home'),('Nagpur','440012','Neuron Hospital'),('Nagpur','440012','Avanti Heart institute'),('Nagpur','440012','Shrikrishna Hrudayalaya'),('Nagpur','440022','G. T. Padole Hospital'),('Nagpur','440012','Criticare Hospital'),('Nagpur','440012','NIMS Hospital'),('Nagpur','440012','Shivganga Hospital'),('Nagpur','440009','Wanjari Hospital'),('Nagpur','440009','Suryodaya Hospital'),('Nagpur','440010','Nandanwar Trauma Center'),('Nagpur','440018','Evershine Hospital'),('Nagpur','440008','Safe Hands Hospital'),('Nagpur','440008','New Era (East End) Hospital'),('Nagpur','440035','Aastha Hospital (Bhandara road)'),('Nagpur','440001','VIIMS Hospital'),('Nagpur','440008','Shri Bhavani Hospital'),('Nagpur','440017','Samarpan Hospital'),('Nagpur','440008','Auro Fracture & Accident Hospital'),('Nagpur','440013','Dr. A. Badar Surgical Nurshing Home'),('Nagpur','440001','Shanti Mohan Hospital'),('Nagpur','440002','Rahate Nursing Home'),('Nagpur','440010','Dande Hospital Hill Road'),('Nagpur','440009','Khalatkar Bal Rugnalaya'),('Nagpur','440013','Carewell Hospital'),('Nagpur','440012','RNH Hospital Pvt. Ltd.'),('Nagpur','440008','Ortho Avenue Hospital'),('Nagpur','440002','Orthonova Hospital & Critical Care Unit'),('Nagpur','440012','Deshmane Hospital'),('Nagpur','440024','Chimalwar Clinic'),('Nagpur','440023','P. D. Hospital'),('Nagpur','440015','Gomashe Nursing Home'),('Nagpur','440024','Mogre Nursing Home & Dental Center'),('Nagpur','440012','Pulse Clinic& Hospital'),('Nagpur','440009','Chaudhari Hospital'),('Nagpur','440008','OM Hospital'),('Nagpur','440009','Shree Ayurveda (Pakwasa)'),('Nagpur','440010','Soni Hospital'),('Nagpur','440012','Tamaskar Clinic'),('Nagpur','440010','Golhar Spine Care & Trauma Research Institute'),('Nagpur','440012','Matru Seva Sangh'),('Nagpur','441002','City Hospital, Kamptee'),('Nagpur','441002','ASHA Hospital, Kamptee'),('Nagpur','441108','Suretech Hospital, Jamtha'),('Nagpur','440012','Ayushman Multispeciality Hsp. Kamptee'),('Nagpur','440023','Nakshatra Hospital, Wadi'),('Nagpur','440016','Axon Hospital, Nagpur'),('Nagpur','441108','Swami Vivekanand Hospital, Khapari'),('Nagpur','440023','Sanjeevani Hospital Godhani'),('Nagpur','441108','COVIDALAYA  Hospital Jamtha (Indumati Gaikwad Patil Hospital)'),('Nagpur','440023','Star Hospital, Wadi'),('Nagpur','441002','Choudhari Hospital, Kamptee'),('Nagpur','441122','Rachana Hospital, Butibori'),('Nagpur','441107','Aditya hospital, Saoner'),('Nagpur','440037','New Life Hospital Dahegaon'),('Nagpur','441106','Kimaya Hospital, Ramtek'),('Nagpur','441002','Kazmi Hospital, Kamptee'),('Nagpur','441002','Roy Hospital Kamptee'),('Nagpur','440019','Shakti Nursing Home, Hingna'),('Nagpur','441001','Mahure Hospital, Kamptee'),('Nagpur','440037','Matoshree Hospital Besa'),('Nagpur','441108','Shri.Sai Multispeciality Hospital, Butibori'),('Nagpur','441203','Zurmure Hospital, Umred'),('Nagpur','441107','Shree Gurukrupa Hospital Saoner'),('Nagpur','440010','G.M.Hospital Bokara'),('Nagpur','440009','Gaikwad Hospital Godhani'),('Nagpur','441107','Jaiswal Hospital Gujarkhedi Saoner'),('Nagpur','441302','Suretech Hospital, Katol'),('Nagpur','440022','Sainath Multispeciality Ghogali'),('Nagpur','440012','Max Care Hospital Khapa Saoner'),('Nagpur','440023','OFAJ Hospital Ambajhari'),('Nagpur','441104','RH Mouda'),('Nagpur','441302','RH Katol'),('Nagpur','441106','SDH Ramtek'),('Nagpur','441001','SDH Kamptee');
/*!40000 ALTER TABLE `data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `database`
--

DROP TABLE IF EXISTS `database`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `database` (
  `Email` varchar(255) NOT NULL,
  `Passwords` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `database`
--

LOCK TABLES `database` WRITE;
/*!40000 ALTER TABLE `database` DISABLE KEYS */;
INSERT INTO `database` VALUES ('admin','admin');
/*!40000 ALTER TABLE `database` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hospital`
--

DROP TABLE IF EXISTS `hospital`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hospital` (
  `namess` varchar(255) NOT NULL,
  `passwords` varchar(255) NOT NULL,
  `oxygen_beds` varchar(255) DEFAULT NULL,
  `ventilator_beds` varchar(255) DEFAULT NULL,
  `normal_bed` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hospital`
--

LOCK TABLES `hospital` WRITE;
/*!40000 ALTER TABLE `hospital` DISABLE KEYS */;
INSERT INTO `hospital` VALUES ('IGGMC','XVlB','270','2','0'),('Alexis Hospital','zgba','4','0','-1'),('Metro Hospital','iCMR','0','0','0'),('Hospital Name','XVlB','O2 Beds','O2 Beds','O2 Beds'),('GMC','zgba','310','2','0'),('Indira Gandhi Rugnalaya Gandhi Nagar','AjWw','45','0','2'),('Ayush Hopital (NMC Sadar)','hTHc','36','0','0'),('Lata Mangeshkar Hospital Hingna (S.G.)','tcuA','10','0','0'),('Lata Mangeshkar Hospital Sitabuldi (S.G.)','xhxK','25','0','14'),('Shalinitai Meghe Hospital (S.G.)','QFDa','20','0','0'),('K.T. Nagar NMC Hospital','FpLS','33','0','0'),('Railway Hospital','jFbc','42','0','0'),('Orange City Hospital','XoEF','11','0','0'),('Kalpavruksha Hospital','fRsW','23','1','0'),('Grace Ortho Hospital','xPLD','1','0','0'),('Swasthyam Superspeciality Hospital','nJOb','26','0','4'),('Sanjeevani Hospital (Laxminagar)','CsNV','10','0','0'),('Viveka Superspeciality Hospital','lgTe','4','0','0'),('Ganga Care Hospital','MaPE','33','2','0'),('Shraman Hospital','ZQle','2','0','0'),('Zenith Hospital','QYhY','4','0','0'),('Central Neurological & Medical Institute','zRyW','3','0','0'),('G.B. Multicare Hospital','JjPj','1','0','0'),('Dande Hospital','zpfR','17','0','0'),('Mure Memorial Hospital','FEgm','20','0','18'),('Kalamkar Multispeciality Hospital','otaF','1','0','0'),('Center Point Hospital','etHs','8','0','0'),('Shri Saikripa Hospital','bZRj','15','0','1'),('Star City Hospital','xAwn','5','0','0'),('Green City Nursing Home','wekr','5','2','0'),('Neuron Hospital','BEmf','6','0','0'),('Avanti Heart institute','dzdc','11','0','0'),('Shrikrishna Hrudayalaya','EkXB','8','0','0'),('G. T. Padole Hospital','AkjQ','5','0','0'),('Criticare Hospital','ZLCt','11','0','0'),('NIMS Hospital','TMtT','2','2','0'),('Shivganga Hospital','CoaN','1','0','0'),('Wanjari Hospital','atyy','2','0','0'),('Suryodaya Hospital','iNKA','9','0','0'),('Borkar Multispeciality Hospital','ReKJ','19','0','0'),('Shravan Hospital','yiXJ','4','0','0'),('Nandanwar Trauma Center','rscc','1','0','0'),('Evershine Hospital','tNsw','1','0','4'),('Sawarkar Multispeciality Hospital','YNsG','1','0','0'),('Central Avenue Critical Care Hospital','Russ','20','0','0'),('Safe Hands Hospital','Vmao','29','1','0'),('New Era (East End) Hospital','zFZB','35','1','0'),('Tarangan Hospital','sbOJ','7','0','10'),('Aastha Hospital (Bhandara road)','iFQG','9','1','0'),('VIIMS Hospital','Zsnw','37','3','0'),('Shri Bhavani Hospital','TKSm','48','2','4'),('Samarpan Hospital','VoiG','8','2','5'),('Rugwani Child Care Center & Hospital','LOpb','18','0','6'),('Gajbe Critical Care & Multispeciality','UOpE','20','0','0'),('Auro Fracture & Accident Hospital','dKup','14','0','0'),('Dr. A. Badar Surgical Nurshing Home','dOMe','2','1','0'),('Shanti Mohan Hospital','RVja','9','0','0'),('Rahate Nursing Home','RzLN','5','0','0'),('Dande Hospital Hill Road','TXYe','11','0','0'),('Khalatkar Bal Rugnalaya','UCWK','4','0','0'),('Carewell Hospital','sXbG','6','0','0'),('RNH Hospital Pvt. Ltd.','yRAO','13','0','0'),('New Life Hospital (Unit of ChhajedHealthcare Pvt. Ltd.)','mBTv','21','0','2'),('Ortho Avenue Hospital','KSJf','2','0','0'),('Orthonova Hospital & Critical Care Unit','jzaL','3','0','0'),('Deshmane Hospital','btZs','6','0','0'),('Akruti Hospital','yMGe','1','0','0'),('Chimalwar Clinic','uDtR','10','0','0'),('P. D. Hospital','zQMD','11','0','4'),('Gomashe Nursing Home','QiYC','3','0','0'),('Mogre Nursing Home & Dental Center','OhgH','11','0','0'),('Pulse Clinic& Hospital','OvgS','2','0','0'),('Chaudhari Hospital','eycJ','15','4','0'),('OM Hospital','PJHY','14','0','0'),('Shree Ayurveda (Pakwasa)','NufN','28','0','0'),('Soni Hospital','jJhh','9','0','0'),('Tamaskar Clinic','jUVR','9','1','0'),('Golhar Spine Care & Trauma Research Institute','uSqf','10','0','0'),('Matru Seva Sangh','gqVM','18','1','0'),('National Cancer Institute Jamtha','kPYV','26','4','0'),('Military Hospital, Kamptee','kURU','13','2','24'),('Life Line Hospital, Kamptee','piFv','36','5','0'),('City Hospital, Kamptee','IZRg','44','0','0'),('ASHA Hospital, Kamptee','BmyA','58','6','0'),('Arch Angel Hspital, Umred','rKCt','18','0','0'),('Suretech Hospital, Jamtha','zkjk','35','0','0'),('Ayushman Multispeciality Hsp. Kamptee','ZIva','18','0','0'),('Nakshatra Hospital, Wadi','BjMk','10','0','0'),('Axon Hospital, Nagpur','XVbW','8','0','7'),('Swami Vivekanand Hospital, Khapari','Gvbq','19','0','0'),('Sanjeevani Hospital Godhani','zgex','3','0','0'),('COVIDALAYA  Hospital Jamtha (Indumati Gaikwad Patil Hospital)','yALB','84','1','35'),('Star Hospital, Wadi','sdjS','15','0','0'),('Choudhari Hospital, Kamptee','Gpng','26','0','6'),('Choudhari Multispeciality Hospital, Butibori','CwFk','13','0','12'),('Shreyas Patil Multispeciality Hospital, Saoner','DifI','19','0','0'),('Rachana Hospital, Butibori','Buuf','19','0','0'),('Yogiraj Swami Sitaramdasji Maharaj Hospital, Ramtek','FMoW','12','0','0'),('Aditya hospital, Saoner','diTs','2','0','15'),('New Life Hospital Dahegaon','kZoQ','2','0','15'),('BAGHE MULTISPECIALITY HOSPITAL','JMqr','3','0','6'),('Shourya Hospital','TICT','17','0','0'),('Kimaya Hospital, Ramtek','ojIY','8','1','0'),('Kazmi Hospital, Kamptee','xyeS','16','0','0'),('Dhanwantari Katol','xZyf','4','0','6'),('Roy Hospital Kamptee','roRO','30','0','2'),('Shakti Nursing Home, Hingna','DMbN','21','0','0'),('Mahure Hospital, Kamptee','DRZn','10','2','0'),('Pande Childcare, Kalmeshwar','PNRW','8','0','9'),('Matoshree Hospital Besa','CJPM','8','0','0'),('Shri.Sai Multispeciality Hospital, Butibori','HDtJ','5','0','0'),('Zurmure Hospital, Umred','mHAY','22','0','18'),('Jawaharlal Neharu Central Hospital, Kandri Kanhan Parseoni','ORsU','12','0','5'),('Shree Gurukrupa Hospital Saoner','fUMA','29','0','20'),('G.M.Hospital Bokara','psVg','4','0','0'),('Gaikwad Hospital Godhani','zHbl','3','0','0'),('Jaiswal Hospital Gujarkhedi Saoner','mYYt','9','0','35'),('Suretech Hospital, Katol','EjVg','12','0','0'),('Sainath Multispeciality Ghogali','wfFb','5','0','0'),('MCH Koradi','bGGc','8','0','0'),('Max Care Hospital Khapa Saoner','nqba','20','0','0'),('OFAJ Hospital Ambajhari','EREu','5','0','30'),('RH Bhiwapur','nUZj','6','0','20'),('RH Mouda','QXmZ','14','0','0'),('RH Katol','OtaR','4','0','0'),('SDH Ramtek','LUtm','14','0','0'),('SDH Kamptee','YgmS','12','0','0');
/*!40000 ALTER TABLE `hospital` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-05-21  1:12:37