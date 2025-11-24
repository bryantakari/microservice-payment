package com.micropayment.userservice.service.impl;

import com.micropayment.userservice.model.entity.Account;
import com.micropayment.userservice.repository.UserRepository;
import com.micropayment.userservice.service.UserService;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
@AllArgsConstructor
public class UserServiceImpl implements UserService {

    private final UserRepository repository;

    @Override
    public Account insertAccount(Account account) {
        return repository.save(account);
    }

    @Override
    public Optional<Account> getAccountByUsername(String username) {
        return repository.findByUsername(username);
    }
}
