package com.micropayment.userservice.common;

import com.micropayment.userservice.config.AppProperties;
import com.micropayment.userservice.config.properties.JwtProperties;
import com.micropayment.userservice.model.dto.JwtDto;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jws;
import io.jsonwebtoken.JwtParser;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import jakarta.annotation.PostConstruct;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

import javax.crypto.SecretKey;
import java.nio.charset.StandardCharsets;
import java.util.Date;
import java.util.Optional;

/**
 * Jwt Helper for any related to jwt.
 */
@Component
@RequiredArgsConstructor
public class JwtHelper {
    private final AppProperties appProperties;
    private JwtProperties jwtProperties;
    private SecretKey key;

    @PostConstruct
    public void init() {
        this.jwtProperties = appProperties.getJwt();
        this.key = Keys.hmacShaKeyFor(jwtProperties.getSecret().getBytes(StandardCharsets.UTF_8));
    }

    public String generateToken(JwtDto dto) {
        return Jwts.builder()
                .subject(dto.getUsername())
                .claim("userId",dto.getUserId())
                .claim("fullname",dto.getFullname())
                .issuedAt(new Date())
                .expiration(
                        new Date((new Date()).getTime() + jwtProperties.getAccessTokenExpiryMillis()))
                .signWith(key)
                .compact();
    }

    public Optional<Claims> validateJwtToken(String token){
        try{
            JwtParser parse = Jwts.parser().verifyWith(key).build();
            Jws<Claims> test = parse.parseSignedClaims(token);

            return Optional.of(test.getPayload());
        }catch (Exception e){
            return Optional.empty();
        }
    }

    public String generateRefreshToken(JwtDto dto) {
        return Jwts.builder()
                .subject(dto.getUsername())
                .claim("userId",dto.getUserId())
                .claim("fullname",dto.getFullname())
                .issuedAt(new Date())
                .expiration(
                        new Date((new Date()).getTime() + jwtProperties.getRefreshTokenExpiryMillis()))
                .signWith(key)
                .compact();
    }


}
