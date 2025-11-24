package com.micropayment.userservice.common.exception.dto;

import com.fasterxml.jackson.annotation.JsonTypeInfo;
import lombok.Getter;
import lombok.experimental.SuperBuilder;

import java.io.Serializable;

/**
 * Represents a detailed error message that provides additional information about an exception.
 * This abstract class is intended to be extended by specific error detail implementations.
 */
@Getter
@JsonTypeInfo(use = JsonTypeInfo.Id.NAME, visible = true)
@SuperBuilder
public abstract class ErrorDetail implements Serializable {

    private String reason;
    private String description;

    protected ErrorDetail(String reason, String description) {
        this.reason = reason;
        this.description = description;
    }

}