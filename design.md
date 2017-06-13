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

Each step will have its own scratch space. MUPPET will allocate the scratch
space and put the path in an environment variable.

## Synchronization between storage systems

It is imperative that files are widely available across several file systems.

File synchronization might take place as a BitTorrent network. Data consumers -
such as THREDDS and MET API - would run customized BitTorrent clients,
configured to retrieve the most recent files from Lustre. BitTorrent might also
help with load distribution in the network.

In order to know which files will be distributed, we need an index of
operational files. Such an index exists today:
[Productstatus](https://github.com/metno/productstatus). The design, speed, and
API of Productstatus could be improved in order to support thousands of queries
per second.

File registration could be facilitated by a new command line utility. After
files are produced, the utility would scan the file, and send its metadata to
Productstatus for indexing.

## Concurrency and parallelism

### Database thread

Responsible for database transactions. Receives pipeline, step, and job object
updates via Go channels, and writes updated information to the database
backend. It must handle database transactions and be ACID compliant, in order
to reduce the amount of potential bugs associated with bad state.

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
