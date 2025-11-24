package com.micropayment.userservice.service.strategy;

import com.micropayment.userservice.model.dto.JwtValidationDto;

/**
 * Token Strategy Pattern interface.
 */
public interface TokenStrategy<T> {
    boolean support(TokenType type);
    String generate(Object subject);
    JwtValidationDto<T> validate(String token);
}
