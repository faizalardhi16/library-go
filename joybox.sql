-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 27, 2022 at 07:09 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `joybox`
--

-- --------------------------------------------------------

--
-- Table structure for table `books`
--

CREATE TABLE `books` (
  `id` varchar(36) NOT NULL,
  `file_name` varchar(100) NOT NULL,
  `title` varchar(100) NOT NULL,
  `penulis` varchar(100) NOT NULL,
  `tahun_terbit` varchar(100) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `books`
--

INSERT INTO `books` (`id`, `file_name`, `title`, `penulis`, `tahun_terbit`, `created_at`, `updated_at`) VALUES
('15eac781-ba78-4cb3-88ef-a555414eaf04', 'images/efd2459f-9df2-407e-82d6-ec9b692d195b-4040237485.jpeg', 'HEHE', 'Faizal', '2020', '2022-11-26 15:44:03.863', '2022-11-27 01:51:41.148'),
('3f83675d-2111-4d5e-ab06-0c07a1cadc4b', 'book.jpg', 'Aku padamu sayang', 'Faizal', '2020', '2022-11-26 15:49:25.633', '2022-11-26 15:49:25.633'),
('b4bc4754-6488-4b4d-8c3f-6dbb145438eb', 'book.jpg', 'The long away', 'Ronan Stewart', '2009', '2022-11-27 09:35:08.043', '2022-11-27 09:35:08.043'),
('c99f4b44-0652-43e6-8630-8969b4be2000', 'book.jpg', 'Aku padamu', 'Faizal', '2020', '2022-11-26 15:47:08.377', '2022-11-26 15:47:08.377'),
('deddb9f2-54c4-48f8-9c71-8669418b8e36', 'book.jpg', 'Bukan aku milikmu', 'Tere Liye', '2020', '2022-11-27 09:34:44.531', '2022-11-27 09:34:44.531');

-- --------------------------------------------------------

--
-- Table structure for table `reservations`
--

CREATE TABLE `reservations` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `booking_date` varchar(100) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `reservations`
--

INSERT INTO `reservations` (`id`, `user_id`, `booking_date`, `created_at`, `updated_at`) VALUES
('013b034b-0f34-443d-b838-922b3753a7a1', '27f65894-b8fa-49f1-8f6c-817b8b36be23', '30-11-2022', '2022-11-27 13:08:44.174', '2022-11-27 13:08:44.174');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `book_id` varchar(36) DEFAULT NULL,
  `reservation_id` varchar(36) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `book_id`, `reservation_id`, `created_at`, `updated_at`) VALUES
('9d485762-2ba1-4690-aa64-ada68b677563', '27f65894-b8fa-49f1-8f6c-817b8b36be23', '15eac781-ba78-4cb3-88ef-a555414eaf04', '013b034b-0f34-443d-b838-922b3753a7a1', '2022-11-27 12:42:42.886', '2022-11-27 13:08:44.181'),
('c55bd4f8-bd17-4370-8481-014d447cf725', '27f65894-b8fa-49f1-8f6c-817b8b36be23', '3f83675d-2111-4d5e-ab06-0c07a1cadc4b', '013b034b-0f34-443d-b838-922b3753a7a1', '2022-11-27 12:42:52.313', '2022-11-27 13:08:44.185');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(100) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `first_name`, `last_name`, `email`, `password`, `role`, `created_at`, `updated_at`) VALUES
('27f65894-b8fa-49f1-8f6c-817b8b36be23', 'Alisson', 'Becker', 'alisson_becker@gmail.com', '$2a$04$67Jj8nWv.9OuFDbbx.L.XOl0CWSDpU4FW5NwKy.L4BLrwF5sU.XXC', 'user', '2022-11-25 14:44:29.237', '2022-11-25 14:44:29.237'),
('438dbe2d-a06c-46a0-98c2-9aa04e833d8f', 'Faizal', 'Ardhi Cahyanto', 'faizal16@gmail.com', '123123', 'user', '2022-11-24 19:48:01.758', '2022-11-24 19:48:01.758'),
('76f43cda-31fe-45ab-8f44-e94efb7fb04c', 'Fifi', 'Aleyda Cahyaningdyah', 'fifialeyda123@gmail.com', '123123', 'user', '2022-11-24 19:52:39.531', '2022-11-24 19:52:39.531'),
('9809c58c-0f21-45de-b18e-b069dd0c3bf5', 'Lucas', 'Tarkowski', 'lucastarkowski@gmail.com', '$2a$04$yOjvLa89ERj.WH3doQXnTuJbO5WtmH80rlrlpxv9jlwpCp0Md8aF.', 'user', '2022-11-25 14:42:59.930', '2022-11-25 14:42:59.930');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_books_id` (`id`);

--
-- Indexes for table `reservations`
--
ALTER TABLE `reservations`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_reservations_id` (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_transactions_id` (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_users_id` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
