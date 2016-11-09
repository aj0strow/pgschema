// The plan package takes a combined DatabaseMatch and produces a list of
// ordered changes to reconcile the prior database state with the new desired
// state. Each change should be one atomic SQL statement.
package plan
