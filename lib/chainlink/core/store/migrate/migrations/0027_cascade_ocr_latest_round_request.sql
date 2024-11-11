-- +goose Up
ALTER TABLE offchainreporting_latest_round_requested
DROP CONSTRAINT offchainreporting_latest_roun_offchainreporting_oracle_spe_fkey,
ADD CONSTRAINT offchainreporting_latest_roun_offchainreporting_oracle_spe_fkey
	FOREIGN KEY (offchainreporting_oracle_spec_id)
	REFERENCES offchainreporting_oracle_specs (id)
	ON DELETE CASCADE
	DEFERRABLE INITIALLY IMMEDIATE;
