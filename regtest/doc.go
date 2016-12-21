// Regression testing package. The folder names are just auto incrementing. Each time pgschema
// breaks in normal use, write a regression test to capture the behavior. Steps to add a new test
// are as follows.
//
// 1. Fire up a new ephemeral schema using the `pgschema/temp` package.
// 2. Create and run a setup SQL script.
// 3. Define the schema in HCL and plan changes.
// 4. Compare actual changes to expected changes. Ignore changes outside schema.
package regtest
