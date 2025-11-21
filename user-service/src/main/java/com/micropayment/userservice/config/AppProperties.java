package com.micropayment.userservice.config;

import com.micropayment.userservice.config.properties.DatasourceProperties;
import com.micropayment.userservice.config.properties.JwtProperties;
import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

/**
 * Properties for application props.
 */
@Getter
@Setter
@Configuration
@ConfigurationProperties(prefix = "app")
public class AppProperties {
    private DatasourceProperties datasource;
    private JwtProperties jwt;
}
