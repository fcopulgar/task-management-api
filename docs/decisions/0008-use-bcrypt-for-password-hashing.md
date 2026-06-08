# 0008 — Usar bcrypt para hash de contrasenas

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

Las contrasenas deben almacenarse de forma segura y nunca en texto plano.

## Decision

Usar bcrypt para hashear y verificar contrasenas.

## Consecuencias

bcrypt es una opcion probada para password hashing. Requiere gestionar costo adecuado y nunca exponer `password_hash`.

## Alternativas consideradas

Hash rapido general y Argon2. Hash rapido se descarta por inseguro para contrasenas; Argon2 queda como alternativa futura si se justifica.

## Relacion con otros ADRs

Se relaciona con autenticacion, cambio de contrasena y gestion de usuarios.
