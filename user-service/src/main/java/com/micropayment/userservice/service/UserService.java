package com.micropayment.userservice.service;


import com.micropayment.userservice.model.entity.Account;

import java.util.Optional;

/**
 * User Service Interface.
 */
public interface UserService {

    Account insertAccount(Account account);

    Optional<Account> getAccountByUsername(String username);
}
