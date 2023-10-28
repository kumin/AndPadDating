CREATE TABLE IF NOT EXISTS `matching` (
  user_id INT NOT NULL,
  partner_id INT NOT NULL,
  is_like TINYINT NOT NULL,

  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, partner_id),
  FOREIGN KEY (user_id) REFERENCES user(id),
  FOREIGN KEY (partner_id) REFERENCES user(id)
);
