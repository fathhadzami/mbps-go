-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 30 Agu 2022 pada 16.48
-- Versi server: 10.4.24-MariaDB
-- Versi PHP: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `webservice`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `attendance`
--

CREATE TABLE `attendance` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `check_in` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `check_in_lng` decimal(10,8) NOT NULL,
  `check_in_lat` decimal(11,8) NOT NULL,
  `check_out` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `check_out_lng` decimal(10,8) NOT NULL,
  `check_out_lat` decimal(11,8) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `department`
--

CREATE TABLE `department` (
  `id` bigint(20) NOT NULL,
  `name` varchar(191) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` varchar(255) NOT NULL,
  `name` int(11) NOT NULL,
  `department_id` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `attendance`
--
ALTER TABLE `attendance`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `department`
--
ALTER TABLE `department`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `department_id` (`department_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `attendance`
--
ALTER TABLE `attendance`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `department`
--
ALTER TABLE `department`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `attendance`
--
ALTER TABLE `attendance`
  ADD CONSTRAINT `attendance_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Ketidakleluasaan untuk tabel `user`
--
ALTER TABLE `user`
  ADD CONSTRAINT `user_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
