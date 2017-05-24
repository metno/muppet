# MUPPET design guidelines

## Abstract

This document attempts to describe the design decisions and problems MUPPET is
trying to solve. The programming paradigm implemented is that of [Flow-based
programming](https://en.wikipedia.org/wiki/Flow-based_programming):

> In computer programming, flow-based programming (FBP) is a programming
> paradigm that defines applications as networks of "black box" processes,
> which exchange data across predefined connections by message passing, where
> the connections are specified externally to the processes. These black box
> processes can be reconnected endlessly to form different applications without
> having to be changed internally. FBP is thus naturally component-oriented.

## Functional programming

As far as possible, avoid data mutation. Instead, manufacture a new copy of the
object, and return it to the caller.

## Output files

In order to have crash-resistant properties, shell scripts should not mutate
data files. If mutation is necessary, the destination file should be copied to
a scratch area first, where the mutation will happen. Then, upon successful
completion, the file can be atomically moved back.

For steps that generate specific files, their scratch space filename and output
filename should be provided by MUPPET itself.

## Synchronization between storage systems

It is imperative that files are being produced on several file systems
simultaneously, without consuming too many resources. Therefore, as each file
is persisted to the storage backend, it should be synchronized across the other
storage systems as well.

How should this be implemented? How should errors be handled? IPFS or similar
mechanisms?

## Concurrency and parallelism

### Database thread

Responsible for database transactions. Receives pipeline, step, and job object
updates via Go channels, and writes updated information to the database backend.

### Cron thread

Maintains a list of steps that should be triggered based on time, and sends
triggered steps to the scheduler.

### Scheduler thread

Maintains the list of all steps. Receives triggered steps, creates jobs in the
pipeline, and persists in the database thread. Database thread sends back to
scheduler thread, which then sends jobs to the executor thread. 

All state changes in the scheduler thread are sent directly to the database
thread for persistance, and then dropped. It is only when the persisted object
comes back from the database thread that further processing is allowed.

### Executor thread

Schedules jobs on executing systems such as the Sun OpenGridEngine, and checks
their exit code, standard output, and standard error.

### API Thread

Receives external requests for product updates, which are sent to the scheduler
thread.
