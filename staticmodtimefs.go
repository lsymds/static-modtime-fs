// Package staticmodtimefs provides a wrapper around an existing filesystem where the ModTime function of a given
// [fs.File] implementation will always return a constant value.
//
// Overriding the default ModTime implementation is mainly useful with embedded and other filesystems which do not
// support native modification times.
package staticmodtimefs
