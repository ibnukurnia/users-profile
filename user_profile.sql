-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 25, 2022 at 04:33 AM
-- Server version: 10.4.20-MariaDB
-- PHP Version: 8.0.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `user_profile`
--

-- --------------------------------------------------------

--
-- Table structure for table `risk_profile`
--

CREATE TABLE `risk_profile` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `bond_percent` float NOT NULL,
  `stock_percent` float NOT NULL,
  `mm_percent` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `risk_profile`
--

INSERT INTO `risk_profile` (`id`, `user_id`, `bond_percent`, `stock_percent`, `mm_percent`) VALUES
(4, 16, 46, 34, 20),
(5, 17, 22, 72, 6),
(6, 18, 22, 72, 6),
(7, 19, 22, 72, 6),
(8, 20, 46, 34, 20),
(9, 21, 46, 34, 20),
(10, 22, 46, 34, 20);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(225) NOT NULL,
  `age` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `age`) VALUES
(16, 'Jon', 45),
(17, 'Budi', 23),
(18, 'Budi', 23),
(19, 'Bambang', 23),
(20, 'Jaja', 46),
(21, 'Jaja', 46),
(22, 'Jaja', 46);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `risk_profile`
--
ALTER TABLE `risk_profile`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_risk_user` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `risk_profile`
--
ALTER TABLE `risk_profile`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `risk_profile`
--
ALTER TABLE `risk_profile`
  ADD CONSTRAINT `fk_risk_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
