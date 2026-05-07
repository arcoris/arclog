/*
   Copyright 2026 The ARCORIS Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// Package predicate defines API-side boolean contracts for deciding whether a
// log entry should continue through a runtime pipeline.
//
// Predicates sit above level thresholds. A level enabler answers whether a
// severity is enabled; a predicate may also inspect entry metadata and already
// constructed fields to express routing, sampling, hook selection, or
// application policy. This package is only the contract layer. It does not
// implement cores, hooks, encoders, writers, registries, or runtime stack
// capture.
//
// Predicate evaluation is expected to be on a hot path. Implementations should
// avoid allocation, reflection, blocking I/O, and avoidable interface churn.
// Predicates must not mutate the supplied Entry or fields. A Predicate value may
// be shared by many loggers and therefore must be safe for concurrent calls
// unless its concrete documentation says otherwise.
//
// Nil is not a valid predicate value. Runtime configuration should use Always
// when no filtering policy is configured and Never when a route, hook, or sink is
// intentionally disabled. Composition helpers do not validate every operand at
// construction time; nil operands panic if evaluation reaches them.
//
// Entry is a small metadata value owned by the caller. The []field.Field slice
// passed to ShouldLog is borrowed for the duration of the call; predicates may
// inspect it but must not retain or mutate it. Caller information may be
// undefined because runtime caller capture can be expensive and may happen after
// predicate checks in some pipelines.
//
// The composition helpers preserve boolean short-circuiting and keep an
// immutable snapshot of their operands. Mutating the slice used to construct a
// composite predicate does not affect future evaluations.
//
// Constructing a multi-operand composite may allocate once to copy operands.
// Evaluation of the built-in composites is allocation-free for the benchmarked
// paths in this repository.
package predicate
