UPDATE db_schema SET config = replace(replace(replace(config::text, '"columnConfigs":', '"columns":'), '"tableConfigs":', '"tables":'), '"schemaConfigs":', '"schemas":')::jsonb;
UPDATE setting SET value = replace(value, '"config":', '"catalog":') WHERE name = 'bb.workspace.schema-template';