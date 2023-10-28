-- Inserting endpoints
INSERT INTO components (type, data) VALUES 
('endpoint', '{"type": "endpoint", "metadata": {"name": "LoginEndpoint", "url": "/login", "method": "POST"}}'),
('endpoint', '{"type": "endpoint", "metadata": {"name": "LogoutEndpoint", "url": "/logout", "method": "GET"}}');

-- Inserting models
INSERT INTO components (type, data) VALUES 
('model', '{"type": "model", "metadata": {"name": "UserModel"}}'),
('model', '{"type": "model", "metadata": {"name": "ProductModel"}}');

-- Inserting actions
INSERT INTO components (type, data) VALUES 
('action', '{"type": "action", "metadata": {"name": "SendEmail"}}'),
('action', '{"type": "action", "metadata": {"name": "ProcessPayment"}}');