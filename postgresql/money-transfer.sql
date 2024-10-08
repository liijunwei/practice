CREATE TABLE accounts (
    id bigserial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    balance DECIMAL(32, 2) NOT NULL CHECK (balance >= 0)
);

insert into accounts (name, balance)
values
	('u1', 100),
	('u2', 100),
	('u3', 100)

-- Create a PL/pgSQL function to transfer money between accounts
CREATE OR REPLACE FUNCTION transfer_funds(sender_id INT, receiver_id INT, amount DECIMAL)
RETURNS VOID AS $$
DECLARE
    sender_balance DECIMAL;
BEGIN
    SELECT balance INTO sender_balance
    FROM accounts
    WHERE id = sender_id;

	-- make sure sender account has sufficient balance to debit
    IF sender_balance >= amount THEN
        -- debit sender account
        UPDATE accounts
        SET balance = balance - amount
        WHERE id = sender_id;

		-- credit receiver account
        UPDATE accounts
        SET balance = balance + amount
        WHERE id = receiver_id;

        RAISE NOTICE 'Transfer successful';
    ELSE
        RAISE EXCEPTION 'Sender does not have sufficient balance for the transfer.';
    END IF;
END;
$$ LANGUAGE plpgsql;

select * from accounts;
SELECT transfer_funds(1, 2, 100.00); select * from accounts;
SELECT transfer_funds(2, 1, 50.00);  select * from accounts;

