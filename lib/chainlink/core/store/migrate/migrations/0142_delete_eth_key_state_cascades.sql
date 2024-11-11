-- +goose Up
ALTER TABLE blockhash_store_specs DROP CONSTRAINT blockhash_store_specs_evm_chain_id_fkey, ADD CONSTRAINT blockhash_store_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE direct_request_specs DROP CONSTRAINT direct_request_specs_evm_chain_id_fkey, ADD CONSTRAINT direct_request_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE eth_txes DROP CONSTRAINT eth_txes_evm_chain_id_fkey, ADD CONSTRAINT eth_txes_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE evm_forwarders DROP CONSTRAINT evm_forwarders_evm_chain_id_fkey, ADD CONSTRAINT evm_forwarders_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE evm_heads DROP CONSTRAINT heads_evm_chain_id_fkey, ADD CONSTRAINT heads_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE evm_key_states DROP CONSTRAINT eth_key_states_evm_chain_id_fkey, ADD CONSTRAINT eth_key_states_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE evm_nodes DROP CONSTRAINT nodes_evm_chain_id_fkey, ADD CONSTRAINT nodes_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE flux_monitor_specs DROP CONSTRAINT flux_monitor_specs_evm_chain_id_fkey, ADD CONSTRAINT flux_monitor_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE keeper_specs DROP CONSTRAINT keeper_specs_evm_chain_id_fkey, ADD CONSTRAINT keeper_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE log_broadcasts DROP CONSTRAINT log_broadcasts_evm_chain_id_fkey, ADD CONSTRAINT log_broadcasts_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE log_broadcasts_pending DROP CONSTRAINT log_broadcasts_pending_evm_chain_id_fkey, ADD CONSTRAINT log_broadcasts_pending_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE log_poller_blocks DROP CONSTRAINT log_poller_blocks_evm_chain_id_fkey, ADD CONSTRAINT log_poller_blocks_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE logs DROP CONSTRAINT logs_evm_chain_id_fkey, ADD CONSTRAINT logs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE ocr_oracle_specs DROP CONSTRAINT offchainreporting_oracle_specs_evm_chain_id_fkey, ADD CONSTRAINT offchainreporting_oracle_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE vrf_specs DROP CONSTRAINT vrf_specs_evm_chain_id_fkey, ADD CONSTRAINT vrf_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) ON DELETE CASCADE DEFERRABLE INITIALLY IMMEDIATE NOT VALID;

-- +goose Down
ALTER TABLE blockhash_store_specs DROP CONSTRAINT blockhash_store_specs_evm_chain_id_fkey, ADD CONSTRAINT blockhash_store_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE direct_request_specs DROP CONSTRAINT direct_request_specs_evm_chain_id_fkey, ADD CONSTRAINT direct_request_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE eth_txes DROP CONSTRAINT eth_txes_evm_chain_id_fkey, ADD CONSTRAINT eth_txes_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE evm_forwarders DROP CONSTRAINT evm_forwarders_evm_chain_id_fkey, ADD CONSTRAINT evm_forwarders_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE evm_heads DROP CONSTRAINT heads_evm_chain_id_fkey, ADD CONSTRAINT heads_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE evm_key_states DROP CONSTRAINT eth_key_states_evm_chain_id_fkey, ADD CONSTRAINT eth_key_states_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE evm_nodes DROP CONSTRAINT nodes_evm_chain_id_fkey, ADD CONSTRAINT nodes_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE flux_monitor_specs DROP CONSTRAINT flux_monitor_specs_evm_chain_id_fkey, ADD CONSTRAINT flux_monitor_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE keeper_specs DROP CONSTRAINT keeper_specs_evm_chain_id_fkey, ADD CONSTRAINT keeper_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE log_broadcasts DROP CONSTRAINT log_broadcasts_evm_chain_id_fkey, ADD CONSTRAINT log_broadcasts_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE log_broadcasts_pending DROP CONSTRAINT log_broadcasts_pending_evm_chain_id_fkey, ADD CONSTRAINT log_broadcasts_pending_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE log_poller_blocks DROP CONSTRAINT log_poller_blocks_evm_chain_id_fkey, ADD CONSTRAINT log_poller_blocks_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE logs DROP CONSTRAINT logs_evm_chain_id_fkey, ADD CONSTRAINT logs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE ocr_oracle_specs DROP CONSTRAINT offchainreporting_oracle_specs_evm_chain_id_fkey, ADD CONSTRAINT offchainreporting_oracle_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;
ALTER TABLE vrf_specs DROP CONSTRAINT vrf_specs_evm_chain_id_fkey, ADD CONSTRAINT vrf_specs_evm_chain_id_fkey FOREIGN KEY (evm_chain_id) REFERENCES evm_chains (id) DEFERRABLE INITIALLY IMMEDIATE NOT VALID;

