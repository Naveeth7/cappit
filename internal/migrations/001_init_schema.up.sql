CREATE TABLE tenants (
                         id UUID PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         api_key VARCHAR(255) UNIQUE NOT NULL,
                         api_secret VARCHAR(255) NOT NULL,
                         created_at TIMESTAMP DEFAULT now(),
                         updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE rate_limits (
                             id UUID PRIMARY KEY,
                             tenant_id UUID REFERENCES tenants(id) ON DELETE CASCADE,
                             route TEXT NOT NULL,
                             method VARCHAR(10) NOT NULL,
                             request_limit INTEGER NOT NULL,
                             window_sec INTEGER NOT NULL,
                             created_at TIMESTAMP DEFAULT now(),
                             updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE webhooks (
                          id UUID PRIMARY KEY,
                          tenant_id UUID REFERENCES tenants(id) ON DELETE CASCADE,
                          url TEXT NOT NULL,
                          event_type VARCHAR(50) NOT NULL,
                          created_at TIMESTAMP DEFAULT now(),
                          updated_at TIMESTAMP DEFAULT now()
);
