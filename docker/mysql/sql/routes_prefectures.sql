CREATE TABLE IF NOT EXISTS `routes_prefectures` (
  `id` int NOT NULL AUTO_INCREMENT,
  `route_id` int NOT NULL,
  `prefecture_id` tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`route_id`) REFERENCES `routes` (`id`),
  FOREIGN KEY (`prefecture_id`) REFERENCES `prefectures` (`id`)
);

INSERT INTO `routes_prefectures` (route_id, prefecture_id) VALUES
  (1, 1),
  (2, 1),
  (3, 1),
  (3, 2),
  (4, 1),
  (5, 1),
  (6, 1),
  (7, 1),
  (8, 1),
  (9, 1),
  (10, 1),
  (11, 1),
  (12, 1),
  (13, 1),
  (14, 1),
  (15, 1),
  (16, 1),
  (17, 3),
  (17, 4),
  (17, 7),
  (17, 9),
  (17, 8),
  (17, 11),
  (17, 13),
  (18, 4),
  (18, 6),
  (19, 4),
  (20, 7),
  (21, 7),
  (21, 15),
  (22, 4),
  (23, 4),
  (24, 3),
  (25, 3),
  (25, 2),
  (26, 2),
  (27, 7),
  (27, 6),
  (28, 7),
  (28, 6),
  (28, 5),
  (28, 2),
  (29, 6),
  (29, 15),
  (30, 6),
  (31, 6),
  (31, 4),
  (32, 6),
  (33, 5),
  (33, 3),
  (34, 3),
  (34, 5),
  (35, 5),
  (35, 3),
  (36, 5),
  (37, 5),
  (37, 2),
  (38, 2),
  (39, 15),
  (39, 6),
  (39, 5),
  (40, 15),
  (41, 2),
  (42, 4),
  (43, 4),
  (43, 3),
  (44, 7),
  (44, 15),
  (45, 3),
  (46, 15),
  (47, 13),
  (48, 13),
  (48, 10),
  (48, 9),
  (48, 8),
  (48, 11),
  (48, 14),
  (49, 14),
  (49, 22),
  (49, 13),
  (50, 14),
  (51, 13),
  (51, 14),
  (52, 14),
  (53, 14),
  (53, 13),
  (54, 13),
  (55, 13),
  (55, 12),
  (56, 12),
  (56, 13),
  (56, 14),
  (57, 14),
  (57, 13),
  (57, 19),
  (57, 20),
  (57, 21),
  (57, 23),
  (58, 13),
  (58, 12),
  (59, 13),
  (60, 13),
  (61, 13),
  (61, 11),
  (61, 8),
  (61, 9),
  (62, 11),
  (62, 10),
  (63, 10),
  (63, 11),
  (63, 13),
  (64, 9),
  (65, 13),
  (65, 12),
  (65, 8),
  (65, 7),
  (65, 4),
  (66, 13),
  (66, 12),
  (66, 8),
  (67, 8),
  (67, 12),
  (67, 13),
  (68, 12),
  (69, 12),
  (70, 12),
  (71, 12),
  (71, 13),
  (72, 13),
  (72, 14),
  (72, 11),
  (72, 12),
  (73, 12),
  (74, 12),
  (75, 10),
  (76, 7),
  (76, 8),
  (77, 8),
  (77, 9),
  (78, 9),
  (78, 10),
  (79, 11),
  (79, 13),
  (79, 14),
  (80, 13),
  (80, 11),
  (81, 11),
  (81, 10),
  (81, 9),
  (81, 12),
  (81, 13),
  (82, 13),
  (83, 13),
  (84, 13),
  (85, 13),
  (85, 12),
  (86, 13),
  (87, 13),
  (87, 11),
  (88, 13),
  (89, 13),
  (90, 13),
  (90, 11),
  (91, 13),
  (91, 12),
  (92, 13),
  (92, 14),
  (93, 13),
  (94, 13),
  (95, 13),
  (95, 12),
  (96, 13),
  (97, 13),
  (97, 14),
  (98, 13),
  (99, 13),
  (100, 13),
  (100, 14),
  (101, 13),
  (102, 14),
  (103, 14),
  (103, 13),
  (104, 14),
  (105, 13),
  (106, 13),
  (106, 11),
  (107, 11),
  (108, 12),
  (109, 13),
  (110, 28),
  (111, 13),
  (111, 11),
  (111, 12),
  (111, 8),
  (112, 12),
  (113, 22),
  (114, 37),
  (114, 36),
  (115, 9),
  (116, 14),
  (117, 12),
  (117, 8),
  (118, 23),
  (118, 21),
  (119, 13),
  (119, 14),
  (119, 22),
  (119, 23),
  (119, 21),
  (119, 25),
  (119, 26),
  (119, 27),
  (119, 28),
  (120, 14),
  (120, 22),
  (121, 22),
  (121, 19),
  (122, 23),
  (122, 22),
  (122, 20),
  (123, 23),
  (124, 21),
  (125, 23),
  (125, 24),
  (125, 26),
  (125, 29),
  (125, 27),
  (126, 24),
  (126, 30),
  (127, 24),
  (128, 24),
  (129, 19),
  (129, 20),
  (130, 20),
  (131, 20),
  (131, 10),
  (131, 15),
  (132, 15),
  (132, 20),
  (133, 15),
  (134, 15),
  (135, 15),
  (135, 10),
  (136, 25),
  (136, 18),
  (136, 17),
  (137, 18),
  (138, 17),
  (139, 16),
  (140, 16),
  (141, 16),
  (141, 21),
  (142, 20),
  (142, 15),
  (143, 23),
  (144, 20),
  (145, 27),
  (146, 27),
  (147, 28),
  (147, 26),
  (148, 26),
  (148, 25),
  (149, 26),
  (149, 27),
  (150, 26),
  (150, 25),
  (151, 25),
  (151, 24),
  (152, 26),
  (153, 26),
  (153, 27),
  (154, 27),
  (154, 28),
  (155, 28),
  (155, 27),
  (156, 27),
  (156, 30),
  (157, 26),
  (157, 29),
  (157, 27),
  (158, 26),
  (159, 29),
  (159, 30),
  (160, 27),
  (161, 24),
  (161, 30),
  (162, 28),
  (163, 28),
  (164, 26),
  (165, 26),
  (165, 18),
  (166, 26),
  (166, 28),
  (167, 28),
  (167, 26),
  (167, 31),
  (167, 32),
  (167, 35),
  (168, 27),
  (169, 27),
  (170, 27),
  (170, 26),
  (171, 27),
  (171, 28),
  (172, 28),
  (173, 27),
  (174, 29),
  (175, 28),
  (175, 27),
  (175, 26),
  (176, 27),
  (176, 30),
  (177, 29),
  (177, 27),
  (178, 44),
  (178, 43),
  (179, 28),
  (179, 27),
  (180, 28),
  (180, 33),
  (181, 33),
  (181, 28),
  (182, 33),
  (183, 33),
  (184, 33),
  (184, 31),
  (185, 33),
  (186, 33),
  (186, 31),
  (187, 31),
  (188, 33),
  (188, 34),
  (189, 34),
  (190, 34),
  (190, 32),
  (191, 32),
  (191, 34),
  (192, 34),
  (193, 34),
  (194, 35),
  (195, 35),
  (195, 32),
  (196, 35),
  (197, 35),
  (197, 28),
  (197, 33),
  (197, 34),
  (197, 40),
  (198, 35),
  (199, 35),
  (200, 33),
  (200, 37),
  (201, 37),
  (201, 38),
  (202, 37),
  (202, 36),
  (202, 39),
  (203, 36),
  (204, 36),
  (205, 38),
  (205, 39),
  (206, 36),
  (207, 40),
  (207, 41),
  (207, 43),
  (207, 46),
  (208, 46),
  (208, 45),
  (208, 44),
  (208, 40),
  (209, 40),
  (209, 44),
  (210, 40),
  (210, 41),
  (211, 41),
  (211, 42),
  (212, 42),
  (212, 41),
  (213, 42),
  (214, 43),
  (214, 45),
  (214, 46),
  (215, 46),
  (215, 45),
  (216, 45),
  (216, 46),
  (217, 46),
  (218, 40),
  (219, 40),
  (220, 45),
  (221, 40),
  (221, 44),
  (222, 41),
  (223, 40),
  (224, 40),
  (225, 43),
  (226, 39),
  (227, 23),
  (228, 3),
  (229, 28),
  (229, 33),
  (229, 31),
  (230, 23),
  (231, 28),
  (232, 40),
  (233, 40),
  (234, 40),
  (235, 23),
  (236, 27),
  (236, 28),
  (237, 28),
  (237, 27),
  (238, 23),
  (239, 23),
  (240, 13),
  (240, 14),
  (241, 14),
  (242, 12),
  (243, 12),
  (149, 26),
  (149, 27),
  (155, 27),
  (155, 28),
  (147, 28),
  (147, 26),
  (247, 13),
  (248, 13),
  (249, 13),
  (249, 11),
  (250, 40),
  (251, 40),
  (251, 41),
  (252, 23),
  (252, 21),
  (252, 25),
  (253, 3),
  (154, 27),
  (154, 28),
  (255, 13),
  (255, 14),
  (256, 23),
  (257, 29),
  (258, 29),
  (258, 26),
  (259, 24),
  (259, 23),
  (260, 24),
  (261, 24),
  (262, 23),
  (263, 23),
  (264, 3),
  (265, 27),
  (265, 30),
  (266, 22),
  (266, 23),
  (267, 13),
  (268, 13),
  (268, 12),
  (269, 40),
  (269, 41),
  (269, 42),
  (269, 43),
  (269, 46),
  (270, 27),
  (270, 29),
  (270, 24),
  (271, 27),
  (271, 29),
  (272, 13),
  (272, 14),
  (273, 13),
  (273, 11),
  (274, 13),
  (274, 14),
  (275, 13),
  (276, 40),
  (277, 13),
  (277, 12),
  (278, 28),
  (279, 28),
  (280, 14),
  (281, 27),
  (282, 13),
  (283, 8),
  (283, 9),
  (284, 13),
  (285, 13),
  (286, 20),
  (286, 15),
  (287, 20),
  (287, 15),
  (146, 27),
  (289, 26),
  (290, 15),
  (290, 17),
  (290, 16),
  (290, 18),
  (290, 25),
  (291, 26),
  (291, 28),
  (291, 31),
  (291, 32),
  (291, 35),
  (292, 13),
  (292, 11),
  (293, 28),
  (293, 33),
  (293, 34),
  (293, 35),
  (294, 23),
  (294, 21),
  (295, 13),
  (295, 14),
  (296, 13),
  (296, 11),
  (297, 13),
  (298, 13),
  (299, 13),
  (299, 14),
  (300, 21),
  (300, 16),
  (301, 15),
  (301, 6),
  (301, 5),
  (302, 26),
  (302, 28),
  (303, 14),
  (304, 11),
  (305, 7),
  (306, 7),
  (307, 13),
  (307, 11),
  (308, 7),
  (309, 13),
  (309, 11),
  (309, 8),
  (309, 9),
  (310, 23),
  (311, 28),
  (312, 28),
  (313, 28),
  (314, 13),
  (314, 12),
  (314, 8),
  (315, 28),
  (316, 20),
  (316, 15),
  (317, 6),
  (317, 5),
  (318, 27),
  (319, 14),
  (320, 13),
  (320, 11),
  (320, 10),
  (320, 9),
  (321, 9),
  (321, 11),
  (321, 10),
  (322, 28),
  (322, 33),
  (322, 31),
  (323, 25),
  (324, 23),
  (325, 13),
  (326, 13),
  (327, 13),
  (327, 11),
  (329, 37),
  (328, 33),
  (330, 23),
  (331, 6),
  (332, 15),
  (333, 27),
  (334, 14),
  (334, 13),
  (335, 13),
  (335, 14),
  (336, 15),
  (337, 10),
  (338, 40),
  (339, 21),
  (339, 23),
  (340, 13),
  (340, 11),
  (340, 9),
  (340, 10),
  (341, 22),
  (342, 13),
  (343, 22),
  (344, 27),
  (344, 28),
  (345, 16),
  (346, 27),
  (346, 29),
  (346, 30),
  (347, 30),
  (347, 27),
  (348, 27),
  (348, 29),
  (349, 27),
  (349, 30),
  (350, 30),
  (350, 27),
  (350, 29),
  (351, 15),
  (352, 13),
  (352, 12),
  (353, 23),
  (354, 22),
  (355, 30),
  (356, 27),
  (356, 30),
  (358, 27),
  (359, 27),
  (360, 13),
  (361, 11),
  (362, 19),
  (362, 20),
  (363, 46),
  (363, 43),
  (364, 22),
  (365, 27),
  (357, 23),
  (366, 27),
  (366, 30),
  (367, 29),
  (368, 21),
  (368, 23),
  (369, 20),
  (370, 23),
  (371, 23),
  (372, 28),
  (372, 33),
  (372, 34),
  (372, 35),
  (372, 40),
  (373, 23),
  (374, 13),
  (374, 11),
  (374, 9),
  (375, 27),
  (376, 27),
  (376, 24),
  (376, 23),
  (377, 27),
  (377, 28),
  (378, 28),
  (378, 27),
  (379, 28),
  (380, 13),
  (380, 14),
  (381, 12),
  (382, 12),
  (383, 24),
  (384, 11),
  (385, 13),
  (385, 12),
  (386, 3),
  (387, 33),
  (387, 34),
  (388, 12),
  (389, 16),
  (390, 27),
  (392, 26),
  (393, 23),
  (394, 27),
  (395, 23),
  (396, 26),
  (397, 10),
  (398, 13),
  (399, 23),
  (400, 27),
  (401, 23),
  (401, 21),
  (402, 21),
  (403, 13),
  (403, 12),
  (404, 27),
  (405, 27),
  (406, 42),
  (407, 27),
  (408, 27),
  (408, 29),
  (409, 22),
  (410, 14),
  (411, 20),
  (412, 15),
  (413, 14),
  (414, 27),
  (415, 23),
  (416, 23),
  (417, 9),
  (417, 8),
  (417, 11),
  (417, 10),
  (417, 12),
  (417, 13),
  (417, 14),
  (417, 22),
  (418, 17),
  (419, 16),
  (420, 27),
  (421, 13),
  (421, 11),
  (421, 9),
  (422, 10),
  (422, 11),
  (422, 13),
  (423, 13),
  (423, 11),
  (423, 10),
  (423, 9),
  (425, 15),
  (426, 24),
  (427, 39),
  (428, 27),
  (429, 29),
  (430, 28),
  (431, 23),
  (432, 40),
  (433, 28),
  (434, 13),
  (435, 9),
  (436, 14),
  (437, 13),
  (438, 10),
  (438, 9),
  (439, 40),
  (440, 27),
  (441, 19),
  (442, 38),
  (443, 37),
  (444, 28),
  (445, 9),
  (446, 23),
  (447, 38),
  (448, 37),
  (449, 40),
  (450, 40),
  (451, 25),
  (452, 37),
  (453, 37),
  (453, 38),
  (454, 27),
  (454, 30),
  (455, 10),
  (456, 30),
  (457, 12),
  (458, 12),
  (459, 8),
  (459, 9),
  (460, 7),
  (461, 7),
  (461, 4),
  (462, 4),
  (463, 8),
  (464, 29),
  (465, 13),
  (466, 38),
  (467, 23),
  (468, 11),
  (469, 27),
  (470, 21),
  (471, 12),
  (472, 13),
  (473, 23),
  (474, 23),
  (475, 29),
  (477, 33),
  (478, 1),
  (479, 33),
  (480, 29),
  (481, 13),
  (482, 27),
  (483, 3),
  (483, 4),
  (484, 4),
  (484, 3),
  (485, 26),
  (486, 26),
  (486, 25),
  (487, 25),
  (488, 1),
  (489, 22),
  (490, 24),
  (491, 10),
  (492, 22),
  (493, 12),
  (494, 8),
  (495, 30),
  (496, 24),
  (497, 24),
  (498, 20),
  (499, 10),
  (500, 24),
  (500, 21),
  (501, 27),
  (502, 28),
  (503, 23),
  (504, 27),
  (504, 26),
  (505, 26),
  (506, 27),
  (507, 23),
  (508, 27),
  (508, 26),
  (509, 38),
  (510, 12),
  (511, 27),
  (511, 30),
  (512, 27),
  (513, 36),
  (513, 39),
  (514, 27),
  (515, 27),
  (515, 28),
  (516, 21),
  (517, 27),
  (517, 28),
  (518, 27),
  (519, 27),
  (520, 24),
  (521, 17),
  (522, 17),
  (523, 27),
  (524, 26),
  (525, 28),
  (526, 28),
  (527, 14),
  (528, 21),
  (529, 2),
  (530, 2),
  (531, 2),
  (531, 1),
  (532, 42),
  (533, 43),
  (534, 33),
  (535, 29),
  (536, 27),
  (537, 14),
  (538, 7),
  (539, 20),
  (540, 23),
  (540, 21),
  (540, 25),
  (541, 22),
  (541, 23),
  (542, 24),
  (542, 30),
  (543, 21),
  (543, 16),
  (544, 23),
  (544, 24),
  (544, 26),
  (544, 29),
  (544, 27),
  (545, 3),
  (546, 14),
  (547, 13),
  (547, 11),
  (547, 14),
  (548, 14),
  (548, 13),
  (549, 22),
  (550, 9),
  (551, 9),
  (552, 26),
  (553, 5),
  (554, 22),
  (555, 40),
  (556, 43);