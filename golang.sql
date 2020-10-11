-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 14, 2020 at 03:42 PM
-- Server version: 10.3.23-MariaDB
-- PHP Version: 7.2.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang`
--
CREATE DATABASE IF NOT EXISTS `golang` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `golang`;

-- --------------------------------------------------------

--
-- Table structure for table `abouts`
--

CREATE TABLE `abouts` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `text` text DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `abouts`
--

INSERT INTO `abouts` (`id`, `title`, `text`, `created_at`, `updated_at`) VALUES
(1, 'What Is Golang ?', 'Go, or Golang, is an open source programming language. It’s statically typed and produces compiled machine code binaries. Developers say that Google\'s Go language is the C for the twenty-first century when it comes to syntax. However, this new programming language includes tooling that allows you to safely use memory, manage objects, collect garbage, and provide static (or strict) typing along with concurrency. \r\n\r\nThe world was first introduced to Go in 2009 thanks to Google’s Rob Pike, Robert Griesemer, and Ken Thompson. The main goal of creating Go was to combine the best features of other programming languages:\r\n\r\nEase of use together with state-of-the-art productivity\r\n\r\nHigh-level efficiency along with static typing\r\n\r\nAdvanced performance for networking and the full use of multi-core power', '2020-07-14 22:28:04', '2020-07-14 22:28:04');

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `text` text DEFAULT NULL,
  `publish` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`id`, `title`, `text`, `publish`, `created_at`, `updated_at`) VALUES
(1, 'Using go/analysis to fix your source code', 'As a recap, the go/analysis package is a Go library that allows you to write linters and checkers for Go source code’s in a fast and standardized way. Usually, people use it to implement linters, which means it shows warnings and errors to you. Running it has no side effects.\r\n\r\nHowever, go/analysis is also able to rewrite your source code. How does it work?\r\n\r\nlook here : https://arslan.io/2020/07/07/using-go-analysis-to-fix-your-source-code/', 'yes', '2020-07-14 22:16:50', '2020-07-14 22:16:50'),
(2, 'Vscode go moves to the go team', 'As the VS Code Go extension grows in popularity and as the ecosystem expands, it requires more maintenance and support. Over the past few years, the Go team has collaborated with the VS Code team to help the Go extension maintainers. The Go team also began a new initiative to improve the tools powering all Go editor extensions, with a focus on supporting the Language Server Protocol with gopls and the Debug Adapter Protocol with Delve. Source : https://blog.golang.org/vscode-go', 'yes', '2020-07-14 22:20:08', '2020-07-14 22:20:08'),
(3, 'Pkg.go.dev is open source!', 'We’re excited to announce that the codebase for pkg.go.dev is now open source.\r\n\r\nThe repository lives at go.googlesource.com/pkgsite and is mirrored to github.com/golang/pkgsite. We will continue using the Go issue tracker to track feedback related to pkg.go.dev. \r\nSource : https://blog.golang.org/pkgsite', 'yes', '2020-07-14 22:21:39', '2020-07-14 22:21:39'),
(4, 'Go and CPU Caches', 'According to Jackie Stewart, a three-time world champion F1 driver, having an understanding of how a car works made him a better pilot. Martin Thompson (the designer of the LMAX Disruptor) applied the concept of mechanical sympathy to programming. In a nutshell, understanding the underlying hardware should make us a better developer when it comes to designing algorithms, data structures, etc.\r\n Source : https://medium.com/@teivah/go-and-cpu-caches-af5d32cc5592', 'yes', '2020-07-14 22:22:42', '2020-07-14 22:22:42'),
(5, 'Optional JSON fields in Go', 'One common kind of data stored in a configuration file is options. In this post, I\'ll talk about some nuances we have to be aware of when storing options in JSON and unmarshaling them to Go.\r\n\r\nSpecifically, the most important difference between options and any other data is that options are often, well... optional. \r\nRead More: https://eli.thegreenplace.net/2020/optional-json-fields-in-go', 'yes', '2020-07-14 22:23:45', '2020-07-14 22:23:45'),
(6, 'Three bugs in the Go MySQL Driver', 'Although GitHub.com is still a Rails monolith, over the past few years we\'ve begun the process of extracting critical functionality from our main application, by rewriting some of the code in Go—mostly addressing the pieces that need to run faster and more reliably than what we can accomplish with Ruby.  \r\nRead more at source : https://github.blog/2020-05-20-three-bugs-in-the-go-mysql-driver/', 'yes', '2020-07-14 22:24:15', '2020-07-14 22:24:15');

-- --------------------------------------------------------

--
-- Table structure for table `contacts`
--

CREATE TABLE `contacts` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  `message` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `contacts`
--

INSERT INTO `contacts` (`id`, `name`, `email`, `subject`, `message`, `created_at`, `updated_at`) VALUES
(1, 'Dicky Pratama', 'dicky@example.com', 'Tanya Tentang Golang', 'Halo, Apakah anda sudah mengerjakan tugas golang? Kalau sudah silahkan hubungi saya melalui email yang saya cantumkan', '2020-07-14 22:29:52', '2020-07-14 22:29:52'),
(2, 'Joko Prabowo', 'joko@prabowo.com', 'Tanya tentang IT', 'Halo GoNews, Saya ingin belajar tentang IT dan Golang. Bisa dibantu?', '2020-07-14 22:30:50', '2020-07-14 22:30:50');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Ardi Gunawan', 'ardi@example.com', '$2a$10$2oqTUNVxFXdxcAtf8ycTSelovjYHx8DWQrKv2Meos8YVqNW6XrxuC', '2020-07-14 22:11:45', '2020-07-14 22:11:45');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `abouts`
--
ALTER TABLE `abouts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `contacts`
--
ALTER TABLE `contacts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `abouts`
--
ALTER TABLE `abouts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `contacts`
--
ALTER TABLE `contacts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
