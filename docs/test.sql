INSERT INTO service (service_name, price) VALUES ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12),
                                                 ('qwerty', 12);







SELECT * FROM service LIMIT 10 OFFSET 10;

SELECT service_id, SUM(money) FROM accounting_report WHERE record_month = 10 AND record_year = 2022 GROUP BY service_id